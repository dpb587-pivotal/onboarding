---
title: Update number of compilation workers in cloud-config
labels:
- cf-deployment
type: story-feature
weight: 11
---

Update number of compilation workers in cloud-config
### What?
GCP creates fresh projects with surprisingly low CPU quotas, which means we generally have to reduce the number of machines we use if we want to deploy CF on a fresh project. [Compilation workers](https://bosh.io/docs/deployment-manifest.html#compilation) temporarily increase your CPU usage beyond what will ultimately be needed by CF once it's up and running, so since we have flexibility on how many we need this is where we have wiggle room to slip in under the barebones quota.

To do this, we edit our **[cloud config](http://bosh.io/docs/cloud-config.html)**, a YAML file that defines the IaaS-specific configuration used by the BOSH Director and all of its deployments. It allows us to separate this configuration out into its own file and keep deployment manifests IaaS-agnostic.

To deploy Cloud Foundry you'll be taking advantage of the [cloud config](http://bosh.io/docs/cloud-config.html) that `bbl` generated for you during `bbl up`.

### How?
Check your CPU quota by navigating to Menu > Compute Engine > Quotas. If you have more than 24 us-east1 CPUs available, then there's no need to change the cloud config but knowing how to do so for the future is a good thing, so read on.

Run `bosh cloud-config > cloud-config.yml`

Then, open the new `cloud-config.yml` file in an editor and locate `compilation: workers:`. If you have a 24 CPU quota on GCP, reduce the number of compilation machines to two; if you have something way higher than 24 CPU leave it alone.

Then, if you edited the file, run `bosh update-cloud-config cloud-config.yml` to apply your changes.

### Expected Result
When you run `bosh -d cf deploy` a couple stories down, it doesn't fail with a reached CPU quota error.

### Resources
[Docs: What is a cloud config?](http://bosh.io/docs/cloud-config.html)
[Forum question: What's the difference between > and >>?](http://askubuntu.com/questions/382793/whats-is-the-difference-between-and-in-shell-command)
[YAML Validator](http://codebeautify.org/yaml-validator)