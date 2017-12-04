---
title: Tear down your CF and Concourse deployments
labels:
- gcp
type: story-feature
weight: 1
---

Tear down your CF and Concourse deployments
### What?
Do you have any idea **[how much all this costs](https://cloud.google.com/products/calculator)?!** Time to clean up after ourselves.

### How?
List BOSH deployments by running `bosh deployments`.
Run `bosh -d <DEPLOYMENT-NAME> delete-deployment` for each of them.