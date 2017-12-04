---
title: Kill a specific instance of an application
labels:
- app-dev
type: story-feature
weight: 11
---

Kill a specific instance of an application
### What?
Sometimes it's clear that a particular instance of your app is having a problem, but it isn't being identified and culled by the [HealthCheck](https://docs.cloudfoundry.org/devguide/deploy-apps/healthchecks.html). This is a good way to take it out yourself.

### How?
Run `watch cf app dora`.

In another buffer, kill a specific instance of dora using `cf curl` (reading the `cf curl` help content may clarify some things).

### Expected Result
You should see the correct instance of your app dying and recovering.

### Resources
[API Docs: Terminate the running app instance at the given index](http://apidocs.cloudfoundry.org/253/apps/terminate_the_running_app_instance_at_the_given_index.html)