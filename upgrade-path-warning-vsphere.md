## ⚠️ Upgrade Path Warning

The following components/apps in **v32.2.0** have versions that are **newer** than those in **v33.0.1** (the next major release on the upgrade path).

This means customers upgrading from **v32.2.0** to **v33.0.1** would experience component/app **downgrades**, which may cause issues.

### Apps

| App | This Release (v32.2.0) | Next Major (v33.0.1) | Impact |
|-----|----------------------------------------|----------------------------------------|--------|
| `cert-exporter` | **2.9.13** | 2.9.12 | ⬇️ Downgrade on upgrade |
| `cert-manager` | **3.9.4** | 3.9.3 | ⬇️ Downgrade on upgrade |
| `etcd-defrag` | **1.2.3** | 1.2.1 | ⬇️ Downgrade on upgrade |
| `etcd-k8s-res-count-exporter` | **1.10.10** | 1.10.9 | ⬇️ Downgrade on upgrade |
| `k8s-audit-metrics` | **0.10.9** | 0.10.8 | ⬇️ Downgrade on upgrade |
| `node-exporter` | **1.20.8** | 1.20.7 | ⬇️ Downgrade on upgrade |
| `security-bundle` | **1.15.0** | 1.13.1 | ⬇️ Downgrade on upgrade |

### Components

| Component | This Release (v32.2.0) | Next Major (v33.0.1) | Impact |
|-----------|----------------------------------------|----------------------------------------|--------|
| `cluster-vsphere` | **3.4.0** | 3.2.0 | ⬇️ Downgrade on upgrade |
| `flatcar` | **4230.2.4** | 4230.2.3 | ⬇️ Downgrade on upgrade |
| `os-tooling` | **1.26.2** | 1.26.1 | ⬇️ Downgrade on upgrade |

> **Note:** This check was triggered because components/apps were automatically bumped to their latest versions.
> If this was intentional, you can ignore this warning and proceed with appropriate documentation.
