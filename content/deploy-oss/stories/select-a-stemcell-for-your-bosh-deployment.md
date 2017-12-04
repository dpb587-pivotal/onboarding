---
title: Select a stemcell for your BOSH deployment
labels:
- cf-deployment
type: story-feature
weight: 12
---

Select a stemcell for your BOSH deployment
### What?
The **[stemcell](http://bosh.cloudfoundry.org/docs/stemcell.html)** is the foundation of every VM the [BOSH Director](https://bosh.io/docs/bosh-components.html#director) deploys. It is a versioned Operating System [image](https://docs.docker.com/engine/userguide/storagedriver/imagesandcontainers/#images-and-layers) wrapped with [IaaS](https://en.wikipedia.org/wiki/Cloud_computing#Infrastructure_as_a_service_.28IaaS.29)-specific packaging.

A typical stemcell contains a bare bones OS skeleton with a few common utilities pre-installed, some configuration files to securely configure the OS by default, and a BOSH Agent. The [BOSH Agent](https://bosh.io/docs/bosh-components.html#agent) is there to listen for instructions from the Director and to carry them out.

Stemcells are distributed as [tarballs](https://bosh.io/docs/build-stemcell.html#tarball-structure). You need to upload a stemcell for your BOSH Director to use when it deploys your Cloud Foundry VMs.

### How?
Available stemcells **[are listed at bosh.io](http://bosh.io/)**. Identify the latest appropriate stemcell for the IaaS  ([Google KVM](https://en.wikipedia.org/wiki/Google_Compute_Engine#Machine_Types)) and OS ([Ubuntu Trusty](https://en.wikipedia.org/wiki/Ubuntu_version_history#Ubuntu_14.04_LTS_.28Trusty_Tahr.29)) you'll be using to deploy your BOSH.

Upload the stemcell by running:
`bosh upload-stemcell https://bosh.io/d/stemcells/bosh-google-kvm-ubuntu-trusty-go_agent?v=VERSION`.

### Resources
[Docs: What is a stemcell?](http://bosh.cloudfoundry.org/docs/stemcell.html)
[Docs: What are light stemcells?](https://bosh.io/docs/build-stemcell.html#light-stemcells)
[Docs: Tarball structure](https://bosh.io/docs/build-stemcell.html#tarball-structure)

### Relevant Repos and Teams
**BOSH:** [cloudfoundry/bosh](https://github.com/cloudfoundry/bosh)
**BOSH:** [cloudfoundry/bosh-agent](https://github.com/cloudfoundry/bosh-agent)