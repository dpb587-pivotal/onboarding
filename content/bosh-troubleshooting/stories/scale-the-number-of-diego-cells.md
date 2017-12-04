---
title: Scale the number of Diego Cells
labels:
- bosh operator
type: story-feature
weight: 1
---

Scale the number of Diego Cells
### What?
With BOSH it is easy to scale deployments. All you need to do is modify the number of instances of that [job](https://bosh.io/docs/jobs.html) in the deployment manifest and redeploy.

### How?
1. Edit the number of Diego Cell instances in the manifest
1. Perform a deploy using `bosh -d DEPLOYMENT_NAME deploy MANIFEST`
1. Observe the new VM appearing by running `bosh vms`

###  Expected Result
You should easily be able to scale the number of Diego Cells up or down. What happens to your apps at that point? Are they redistributed as soon as there is a new cell or do you have to scale the app to trigger a relocation?
cf-deployment installs [`cfdot`](https://github.com/cloudfoundry/cfdot) on the diego cells, which you can use to interrogate Diego. The `watch` and `tail` bash commands will also be your friends during this investigation.

### Resources
[Docs: BOSH Deploying, a step-by-step walk through](http://bosh.io/docs/deploying-step-by-step.html)
[Docs: BOSH Deployment Manifest docs](http://bosh.io/docs/deployment-manifest.html)
[Docs: Creating a new BOSH VM](http://bosh.io/docs/bosh-components.html#create-vm)
[YAML Validator](http://codebeautify.org/yaml-validator)