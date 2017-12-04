---
title: Follow the Concourse "Hello, world!" tutorial
labels:
- concourse
type: story-feature
weight: 4
---

Follow the Concourse "Hello, world!" tutorial
### What?
To get started with Concourse, you'll use [their official "Hello, world!" tutorial](http://concourse.ci/hello-world.html) to set up two very simple pipelines.

The first has a single [job](http://concourse.ci/concepts.html#jobs) that runs a single [task](http://concourse.ci/concepts.html#tasks) and can only be kicked off manually (i.e. via the CLI or the GUI).

The second also has a single job, but it is [triggered](http://concourse.ci/get-step.html#trigger) by a `time` [resource](http://concourse.ci/concepts.html#resources) that will cause the job to run every minute.

### How?
Just follow along with [the tutorial](http://concourse.ci/hello-world.html)! As you go, take the time to learn about:
1. [Configuring Jobs](http://concourse.ci/configuring-jobs.html)
1. [Configuring Resources](http://concourse.ci/configuring-resources.html)

### Expected Result
A pipeline with a job that runs every minute, plus a basic understanding of these [core Concourse concepts](http://concourse.ci/concepts.html). You'll be setting up your own soon!

### Troubleshooting
Have you clicked "Login"? You won't see your pipeline or jobs in your browser until you do so.

### Resources
[Docs: Pipeline Mechanics](http://concourse.ci/pipeline-mechanics.html)