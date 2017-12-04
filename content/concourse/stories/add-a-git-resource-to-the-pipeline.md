---
title: Add a Git resource to the pipeline
labels:
- concourse
type: story-feature
weight: 4
---

Add a Git resource to the pipeline
### What?
A  **[resource](http://concourse.ci/concepts.html#resources)** is any entity that can be checked for new versions, pulled down at a specific version, and/or pushed up to idempotently create new versions. A few of the usual suspects are [listed here](http://concourse.ci/resource-types.html). They include a number of ideas that you're used to thinking of in terms of versioning (`git resource`, `git-release resource`, `tracker resource`, etc.), but also a few that you might not be, like `time`.

The [`git resource`](https://github.com/concourse/git-resource) tracks the commits in a git repository and, though I have no numbers on this, I expect it's the most commonly used Concourse resource. And you can have one of your very own!

### How?
Create a repo to house your Concourse code (`pipeline.yml` etc.) and set up a [git resource](https://github.com/concourse/git-resource) that fetches it. It will be useful to have access to this once we add tasks and a Dockerfile that you'll want to reference in your pipeline. Move your inline task to a yaml file and either a Bash or Ruby script in that repo.

### Expected Result
A one job, one resource pipeline that is always green.

### Resources
[Pipeline Mechanics](http://concourse.ci/pipeline-mechanics.html)
[Common Concourse "resources"](http://concourse.ci/concepts.html#resources)