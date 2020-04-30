# Announcing Giant Swarm App Platform
We are pleased to announce that the Giant Swarm App Platform is generally available.

In the past few months, you may have noticed that we have been working on this. We even answered some questions we received as a result of the preview. This has helped us shape the platform and bring it to where it is today.

There are two ways to benefit:

* One, install the apps from the Giant Swarm Catalog on your clusters. Some apps we are working towards fully supporting are NGINX Ingress Controller, Kong, and Prometheus Operator.
* Two, create your own App Catalogs and build a pipeline of your own apps into your clusters.

Below, you can find more details on the multiple catalogs in the current App Platform.


## 1. **The Giant Swarm Catalog**

The Giant Swarm Catalog contains our stable, fully managed apps, with SLAs (e.g. the NGINX Ingress Controller). 

It also contains apps we intend to develop towards that level of commitment (e.g. Kong, EFK).

Please note, not all apps are at the same level of maturity. The fact that they appear in the Giant Swarm Catalog does not make them automatically fully managed by us. Please contact us if you have any questions with regards to a specific app.


![The Giant Swarm Catalog](https://p80.f1.n0.cdn.getcloudapp.com/items/lluDyJ14/Image%202020-04-22%20at%209.58.56%20AM.png?v=a356238b6f1a34f5840d3609c743f808)

## 2. The Playground Catalog

The Playground Catalog is our go-to place to create and try things out. 

Mainly, it contains apps that a Solution Engineer has added to help you with a specific issue. 

Additionally, you will find some apps created for non-commercial purposes (e.g. for a blog post or workshop).


![The Playground Catalog](https://p80.f1.n0.cdn.getcloudapp.com/items/z8uxX8wL/Image%202020-04-22%20at%2010.03.00%20AM.png?v=1439b2120f9203b7e4ba13ce26b83af3)


Bear in mind that we do NOT support Playground apps and they aren’t a priority in terms of our workload. This means they might not ever make it into the Giant Swarm Catalog. What you will get is an app at a basic maturity level at a specific point in time.


## 3. Your Enterprise App Catalog

Finally, the App Platform infrastructure enables you to create an enterprise App Catalog with your own apps. This lets you provide a standards-compliant and easy-to-use set of apps to your engineers.

It’s already possible for you to add your own catalog(s) via the Kubernetes API in your Control Plane. We are working on documentation for how to do this. In the meantime, you can request access from your Solution Engineer.


----------

We encourage you to check out the app catalogs and try some of the apps there. You can also work with your Solution Engineer to create your own enterprise app catalog. As always, feedback is welcome.

We look forward to enabling you to build the cloud-native stack that is most useful to you with the Giant Swarm App Platform.
