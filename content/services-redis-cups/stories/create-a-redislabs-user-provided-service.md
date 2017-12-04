---
title: Create a RedisLabs user-provided service
labels:
- app-dev
- services
type: story-feature
weight: 1
---

Create a RedisLabs user-provided service
### What?
Cloud Foundry offers a [marketplace of services](https://docs.run.pivotal.io/marketplace/) from which users can provision reserved resources on-demand. Think of a service as a factory that delivers service instances. Examples include databases on a shared or dedicated server or accounts on a Software as a Service (SaaS) application.

As an alternative, you can also provide your own service instances in the form of [User-Provided Service Instances](https://docs.cloudfoundry.org/devguide/services/user-provided.html), which enable developers to use services that are not available in their Cloud Foundry marketplace.

Let's create one backed by [RedisLabs](https://redislabs.com/).

### How?
1. Create an account with [RedisLabs](https://redislabs.com/) and generate a New Redis Subscription (select the free 30MB option).
1. As part of setting up that Redis database, set the password (by default there is no password).
1. Run `cf cups redis -p "host, port, password"` and it will prompt you for those attributes. Fill them with the endpoint and port RedisLabs provides, and the password you used when setting up your database.

### Expected Result
I can see my new user-provided service in `cf services`.

### Resources
[Creating a user-provided service](https://docs.cloudfoundry.org/devguide/services/user-provided.html)

### Relevant Repos and Teams
**CAPI:** https://github.com/cloudfoundry/cloud_controller_ng