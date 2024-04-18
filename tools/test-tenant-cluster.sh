#!/bin/bash

installation=$1
cluster=$2

read -p "Installation: $installation, Cluster: $cluster.
Please ensure you have a kubeconfig for the workload cluster before you continue.
Press enter to continue, ctrl+c to abort"

context=gs-$installation-$cluster-clientcert

# Wait for at least one node to become ready
count=0
while [ $count -lt 1 ]
do
  count=$(kubectl --context=$context get no -L role|grep worker|grep  -v "NotReady"|wc -l)
  if [ $count -lt 1 ]
  then
    sleep 5
    echo "Found $count ready nodes"
  fi
done

# Ingress

base="`opsctl show installation -i $installation| yq .Base`"

echo "Deploying helloworld app"

echo "
apiVersion: v1
kind: ConfigMap
metadata:
  name: loadtest-app-user-values
  namespace: ${cluster}
data:
  values: |
    replicaCount: 2
    autoscaling:
      enabled: false
    ingress:
      enabled: true
      paths:
        - /
      hosts:
        - helloworld.${cluster}.k8s.$base
---
apiVersion: application.giantswarm.io/v1alpha1
kind: App
metadata:
  name: loadtest-app
  namespace: ${cluster}
spec:
  catalog: default
  config:
    configMap:
      name: ${cluster}-cluster-values
      namespace: ${cluster}
  install: {}
  kubeConfig:
    context:
      name: ${cluster}-kubeconfig
    inCluster: false
    secret:
      name: ${cluster}-kubeconfig
      namespace: ${cluster}
  name: loadtest-app
  namespace: default
  userConfig:
    configMap:
      name: loadtest-app-user-values
      namespace: ${cluster}
  version: 0.5.0
" | kubectl --context gs-${installation} apply -f -

echo "Waiting for helloworld app to be deployed"

status=""
while [ "$status" != "deployed" ]
do
  if [ "$status" != "" ]
  then
    echo "Waiting for app status to be 'deployed', it is '$status'"
    sleep 5
  fi
  status="$(kubectl --context gs-${installation} -n ${cluster} get app loadtest-app  -o yaml|yq .status.release.status)"
done

echo "Helloworld app is deployed"

count=0
while [ $count -ne 1 ]
do
  count=$(kubectl --context=$context get deploy loadtest-app|grep "2/2"|wc -l)
  if [ $count -ne 1 ]
  then
    echo "Helloworld app is not up yet."
    sleep 5
  fi
done

echo "Hello world app is ready"

url=helloworld.$cluster.k8s.$base

code=""
while [ "$code" != "200" ]
do
  if [ "$code" != "" ]
  then
    echo "Wasn't able to access helloworld app via Ingress http://$url (code=$code)"
    sleep 5
  fi
  code=$(curl http://$url -o /dev/null -I -w "%{http_code}" -s)
done

echo "Hello world app reachable via ingress"

echo "Checking networking"
pods=$(kubectl --context=$context get po -l app.kubernetes.io/name=loadtest-app -o wide| awk '{print $1 "|" $6}'|grep -v NAME)

first_pod_name=$(echo "$pods" | head -n 1| cut -d"|" -f1)
first_pod_ip=$(echo "$pods" | head -n 1| cut -d"|" -f2)
second_pod_name=$(echo "$pods" | tail -n 1| cut -d"|" -f1)
second_pod_ip=$(echo "$pods" | tail -n 1| cut -d"|" -f2)

status=$(kubectl --context=$context exec -ti $first_pod_name -- wget $second_pod_ip:8080 -S -q -O /dev/null 2>&1|grep HTTP|awk '{print $2}')
if [ "$status" != "200" ]
then
  echo "Can't connect from pod $first_pod_name to $second_pod_ip"
  exit 1
fi
status=$(kubectl --context=$context exec -ti $second_pod_name -- wget $first_pod_ip:8080 -S -q -O /dev/null 2>&1|grep HTTP|awk '{print $2}')
if [ "$status" != "200" ]
then
  echo "Can't connect from pod $second_pod_name to $first_pod_ip"
  exit 1
fi

echo "Networking is ok"

echo "Testing PVC"

classes="$(kubectl --context=$context get storageclass -o yaml| yq '.items[].metadata.name')"

i=0
for class in $classes
do
  echo "Testing PVC for storageclass $class"

  echo "apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: myclaim-$i
spec:
  storageClassName: $class
  accessModes:
    - ReadWriteOnce
  volumeMode: Filesystem
  resources:
    requests:
      storage: 8Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: pvc-test-$i
  namespace: default
spec:
  containers:
  - image: quay.io/giantswarm/helloworld:latest
    name: helloworld
    volumeMounts:
    - name: mypd
      mountPath: /mnt
    securityContext:
      allowPrivilegeEscalation: false
      capabilities:
        drop:
        - ALL
      readOnlyRootFilesystem: true
  securityContext:
    runAsGroup: 33
    runAsUser: 33
    runAsNonRoot: true
    seccompProfile:
      type: RuntimeDefault
  volumes:
  - name: mypd
    persistentVolumeClaim:
      claimName: myclaim-$i
" | kubectl --context=$context apply -f - >/dev/null

  count=0
  while [ $count -eq 0 ]
  do
    count=$(kubectl --context=$context get pvc myclaim-$i| grep Bound | wc -l)
    if [ $count -eq 0 ]
    then
      echo "Waiting for PVC to be Bound"
      sleep 5
    fi
  done

  echo "PVC for storageclass $class is Bound"

  kubectl --context=$context delete po pvc-test-$i --wait=false
  kubectl --context=$context delete pvc myclaim-$i --wait=false

  i=$((i+1))
done


echo "Testing PVC for default storageclass"

echo "apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: myclaim
spec:
  accessModes:
    - ReadWriteOnce
  volumeMode: Filesystem
  resources:
    requests:
      storage: 8Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: pvc-test
  namespace: default
spec:
  containers:
  - image: quay.io/giantswarm/helloworld:latest
    name: helloworld
    volumeMounts:
    - name: mypd
      mountPath: /mnt
    securityContext:
      allowPrivilegeEscalation: false
      capabilities:
        drop:
        - ALL
      readOnlyRootFilesystem: true
  securityContext:
    runAsGroup: 33
    runAsUser: 33
    runAsNonRoot: true
    seccompProfile:
      type: RuntimeDefault
  volumes:
  - name: mypd
    persistentVolumeClaim:
      claimName: myclaim
" | kubectl --context=$context apply -f - >/dev/null

count=0
while [ $count -eq 0 ]
do
  count=$(kubectl --context=$context get pvc myclaim| grep Bound | wc -l)
  if [ $count -eq 0 ]
  then
    echo "Waiting for PVC to be Bound"
    sleep 5
  fi
done

echo "PVC for default storageclass is Bound"

kubectl --context=$context delete po pvc-test --wait=false
kubectl --context=$context delete pvc myclaim --wait=false
