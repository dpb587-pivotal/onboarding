---
title: Create a pipeline with a single job
labels:
- concourse
type: story-feature
weight: 3
---

Create a pipeline with a single job
### What?
Every great Concourse behemoth of the 21st century started in the same place as yours is about to: with an empty `pipeline.yml`. You can kick it off with as little as a single, manually-triggered job.

### How?
1. Set up a `jobs:` section.
1. Define a **[job](https://concourse.ci/concepts.html#jobs)** with a name and a [plan](http://concourse.ci/build-plans.html).
1. Give the job a task that you define inline, in the pipeline itself (this is good to know how to do to test new configurations quickly).
1. "fly" your pipeline using the `fly` CLI.

### Expected Result
A job that always goes green. How idyllic.

### Resources
[Configuring a Job](http://concourse.ci/configuring-jobs.html)
[Build Plans](http://concourse.ci/build-plans.html)
[Pipeline Mechanics](http://concourse.ci/pipeline-mechanics.html)