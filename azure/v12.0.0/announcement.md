# :zap: Giant Swarm Release 12.0.0 for Azure :zap:

This is the first Giant Swarm Azure release which includes **Kubernetes v1.17** and changes **NGINX Ingress Controller App** from being ** pre-installed to optional** component on Azure. This allows NGINX App installations to be managed independently from the base platform lifecycle.

**Important Warning** 
During master upgrade from 11.4.0 to 12.0.0, within the time frame of 30 second we had noticed a spike in requests failures. This is most likely caused by Azure CNI upgrade and despite our great efforts, we had not found a solution to maintain upgrade path and avoid this disturbance. Please keep this in mind when scheduling an upgrade window, and contact your SE if you have further questions.

Many core components and default apps have been updated for improved reliability and observability. Further details about changes to individual components can be found in the [release notes](https://github.com/giantswarm/releases/blob/master/azure/v12.0.0).

---
Please let <!subteam^S0GSG846L> know if you have any feedback or questions by replying to this announcement in a thread.
