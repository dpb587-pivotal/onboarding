---
title: Stand up Concourse locally
labels:
- concourse
type: story-feature
weight: 2
---

Stand up Concourse locally
### What?
Concourse is a Pivotal-sponsored, pipeline-based, continuous integration and deployment (CI/CD) system.

The "pipelines" are a collection of [three core concepts (jobs, tasks, and resources)](http://concourse.ci/concepts.html) that you'll learn more about in upcoming stories. While CI may call to mind test automation, Pivotal teams use it for so much more than that. Take a stroll around the office and check out the jobs up on the CI screens to get a general idea of how broadly we use it to automate all that is automate-able.

### How?
**[Follow these instructions](http://concourse.ci/concourse-lite.html)** to spin up a local Concourse VM with VirtualBox.

### Expected Result
![Lonely Concourse, no pipelines](http://engineering.pivotal.io/images/concourse-000/no_pipelines.png)

### Resources
[Docs: Concourse Architecture Overview](https://concourse.ci/architecture.html)
[Video: All About Concourse for Continuous Integration (video)](https://blog.pivotal.io/pivotal-perspectives/features/all-about-concourse-for-continuous-integration)

### Relevant Repos and Teams
**Concourse:** [concourse/concourse](https://github.com/concourse/concourse)