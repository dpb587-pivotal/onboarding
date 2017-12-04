---
title: Install Google Cloud SDK
labels:
- bbl
type: story-feature
weight: 2
---

Install Google Cloud SDK
### What?
You're going to deploy your BOSH Director to [Google Cloud Platform (GCP)](https://cloud.google.com). Although [Amazon Web Services (AWS)](https://aws.amazon.com/what-is-aws/) was historically Pivotal Cloud Foundry's go-to IAAS, we've decided to migrate as much as possible to  GCP. GCP is significantly cheaper than AWS and managing costs is significantly easier—no more reserved instances! We also have an excellent relationship with GCP PMs and Engineers and can escalate issues directly through our contacts.

The `gcloud` CLI manages authentication, local configuration, developer workflow, and interactions with the Cloud Platform APIs. Download and install the [Google SDK](https://cloud.google.com/sdk) to get `gcloud`, as you'll need it to `bbl up`.

This is another "fun-sized" story (you know you'll miss them once you hit the 6-hour ones....)

### How?
Follow **[these instructions](https://cloud.google.com/sdk/downloads)** to install Google Cloud SDK on your workstation.

Once you've done that, run `gcloud init` in your terminal.

You should have received an invitation to join a GCP project from your facilitator or ask+cf@pivotal.io. Use your @pivotal.io email address and the name of your assigned project to configure the tool (the project name is in the GCP dashboard menu bar and the URL).

### Expected Result
Run `gcloud version` => "Google Cloud SDK 149.0.0 ..." or higher

Run `gcloud config list` =>
```
[core]
account = your-email@pivotal.io
disable_usage_reporting = False
project = PROJECT_ID

Your active configuration is: [PROJECT_ID]
```

### Resources
[Docs: Comparison GCP <> AWS](https://cloud.google.com/docs/compare/aws/)
[Internal Doc: The Great GCP Migration of 2016](https://docs.google.com/document/d/1ze6znVK32UlpsmGXHmMn4ZAAAvuUCK0yVPXhXH74-n0/edit#heading=h.x5ivrrqjuddi)