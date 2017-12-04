---
title: Deploy Concourse with BOSH
labels:
- concourse
type: story-feature
weight: 1
---

Deploy Concourse with BOSH
### What?
[Concourse](http://concourse.ci/) is a Pivotal-sponsored, pipeline-based, continuous integration and deployment (CI/CD) system. The "pipelines" are a collection of [three core concepts (jobs, tasks, and resources)](http://concourse.ci/concepts.html) that you'll learn more about in upcoming stories. While CI may call to mind test automation, Pivotal teams use it for so much more than that. Take a stroll around the office and check out the jobs up on the CI screens to get a general idea of how broadly we use it to automate all that is automate-able.

### How?
You already have a Bosh Director, a cloud-config, and at least one stemcell, so to spin up a full Concourse deployment all you have to do is prepare your deployment manifest. You can follow **[these instructions](https://github.com/cloudfoundry/bosh-bootloader/blob/master/docs/concourse-gcp.md)** to set it up and Bosh deploy, but as per the yooj there's a small set of additional instructions...

**Pre-deploy**
`bbl` only supports one load balancer at a time so in lieu of depriving our CF deployment of its lb, we're going to have Concourse go without. To do this:
1. Omit `tls_bind_port`, `tls_cert`, and `tls_key` attributes from your Concourse manifest.
1. Replace them with `bind_port: 80`

Also, register a new domain on Freenom. Don't do any configuration yet.

**Post-deploy**
 Once `cf -d concourse deploy concourse.yml` is finished running, reserve a static IP address through GCP and attach it to your web VM (you can find its VM CID by running `bosh -d concourse vms`).

Then, on Freenom go to Services > My Domains > Manage Domain > Manage Freenom DNS and create an A record that targets the static IP you just attached to your Concourse web VM.

### Troubleshooting
1. Did you remember to include `http://` in the domain provided as the `external_url` of the `atc` job?
1. Double-check that you're visiting the `http://` address in your browser.
1. Try `bosh ssh`ing into your web VM and curling localhost. It should return the html for your pipeline-less Concourse webpage. If it does, then your problem is with routing/DNS, not with Concourse itself.
1. If you hit an IP quota, go to Home > IAM & Admin > Quotas in your GCP dashboard and click the ✏  button. This will take you to a form where you can request an increased quota.
1. If this ends up being frustrating in a "please, please don't make me do another minute" kind of way, no sweat. Don't waste time on a story you're not getting anything out of, just _Choose Your Own Adventure_ your way out of it by **[spinning up a local VM with `bosh create-env` and VirtualBox](http://concourse.ci/single-page.html#concourse-lite)**. Either way, same result →

### Expected Result
![Lonely Concourse, no pipelines](http://engineering.pivotal.io/images/concourse-000/no_pipelines.png)

### Resources
[How to setup A, MX, CNAME or other records for a domain name on Freenom](https://my.freenom.com/knowledgebase.php?action=displayarticle&id=4)
[Terraform + Google Cloud Platform docs](https://www.terraform.io/docs/providers/google/index.html)
[Concourse Architecture Overview](https://concourse.ci/architecture.html)
[All About Concourse for Continuous Integration (video)](https://blog.pivotal.io/pivotal-perspectives/features/all-about-concourse-for-continuous-integration)
[BOSH 2.0: The Evolution - YouTube (video)](https://www.youtube.com/watch?v=Q5uvoL1OqSw)
[YAML Validator](http://codebeautify.org/yaml-validator)

### Relevant Repos and Teams
**Concourse:** [concourse/concourse](https://github.com/concourse/concourse)