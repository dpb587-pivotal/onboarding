---
title: Add a docker-image resource to the pipeline
labels:
- concourse
tasks:
- Create DockerFile
type: story-feature
weight: 5
---

Add a docker-image resource to the pipeline
### What?
Another popular Concourse resource is the **[`docker-image` resource](https://github.com/concourse/docker-image-resource)**. It's useful for providing a reproducible environment for your tests and builds with all of the dependencies a growing pipeline needs to succeed.

### How?
1. Write a simple Dockerfile.
1. Upload it to the DockerHub registry (if you do not have an account, create one at this point).
1. Set up the `docker-image` resource to fetch it.
1. Use it as the [image resource](http://concourse.ci/single-page.html#image_resource) in your task.

### Resources
[Docker-image resource repo](https://github.com/concourse/docker-image-resource)
[Get Started with Docker](https://docs.docker.com/get-started/)