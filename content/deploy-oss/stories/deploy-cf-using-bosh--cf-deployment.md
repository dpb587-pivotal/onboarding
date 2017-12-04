---
title: Deploy CF using BOSH + cf-deployment
labels:
- cf-deployment
type: story-feature
weight: 13
---

Deploy CF using BOSH + cf-deployment
### What?
[Cf-deployment](https://github.com/cloudfoundry/cf-deployment/blob/master/README.md#purpose) is another bleeding edge technology you'll be using during your CF adventure. It's not quite ready for use in production environments, but it's the future so it makes more sense to learn about this than about something that's about to be deprecated.

The standard cf-deployment is fairly hefty, so we're going to scale ours down to a more reasonable size using an [operations file](https://bosh.io/docs/cli-ops-files.html) that reduces the number of VMs we deploy. Using an operations file to make small adjustments like this means we can edit our deployment without changing the core `cf-deployment.yml` manifest.

### How?
Clone the **[cf-deployment repo](https://github.com/cloudfoundry/cf-deployment)** to your workstation, then follow [these instructions to create a BOSH deployment manifest](https://github.com/cloudfoundry/cf-deployment/blob/master/deployment-guide.md#deploy-cf) (ignore the stemcell step, you'll use the one you uploaded in the last story).

Add `-o operations/scale-to-one-az.yml` to the deploy command.

If you want to speed up the deploy, you can also add `-o operations/use-compiled-releases.yml`, to skip compiling all of your releases.

### Expected Result
Wait for the BOSH deploy to complete, then run `bosh vms`. All of them should have a status of "running".

### Resources
[Docs: What is a BOSH release?](https://bosh.io/docs/release.html)
[Github README: Why are we replacing cf-release with cf-deployment?](https://github.com/cloudfoundry/cf-deployment/blob/master/README.md#purpose)