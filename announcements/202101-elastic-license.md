**tl;dr Elastic is changing Elasticsearch and Kibana**’**s license. Amazon will step up to continue providing it open source. There is likely nothing to change or worry about.**

ℹ️ **Elastic’s decision to change Elasticsearch and Kibana’s license**

You might have heard of [Elastic’s decision](https://www.elastic.co/blog/licensing-change) to change Elasticsearch and Kibana’s license, which our Managed EFK uses under the Amazon-supported Open Distro for Elasticsearch.

By their next release (v7.11), Elastic is changing the license from Apache 2.0 to Server Side Public License (SSPL). Under SSPL, organizations using Elasticsearch and Kibana will be at risk of being forced to release their intellectual property.

[Amazon decided](https://aws.amazon.com/blogs/opensource/stepping-up-for-a-truly-open-source-elasticsearch/) to create and maintain an Apache 2.0 fork of Elasticsearch and Kibana.

**What this means for you:**

- Your current EFK is under Apache 2.0. You are not at risk of having to release your intellectual property.
- We will continue to provide a Managed EFK, using Open Distro, which Amazon plans to create and maintain under the Apache 2.0 license. You can continue to use this.

There is nothing to worry about. We will continue to monitor developments for you.

P.S. We are developing a [Managed Loki](https://github.com/giantswarm/loki-app) to provide as a logging solution. Are you interested? Let's talk!