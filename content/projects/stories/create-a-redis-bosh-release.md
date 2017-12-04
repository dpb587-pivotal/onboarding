---
title: Create a Redis BOSH Release
labels:
- bosh
type: story-feature
weight: 1
---

Create a Redis BOSH Release
### What?
Okay, the training wheels are coming **off**. This BOSH Release is gonna be you + the same docs anyone else uses ("whhaaaaaat?" "I know, they grow up so fast!"). But first, a little preamble:

A BOSH release is a versioned collection of configuration properties, configuration templates, start up scripts, source code, binary artifacts, and anything else required to *build and deploy software in a reproducible way*. Each BOSH VM starts with a [stemcell](http://bosh.io/docs/stemcell.html), and then has utility added to it in the form of releases. Your Redis release might include start-up and shutdown scripts for the redis-server, a tarball with Redis source code obtained from the Redis official website, and configuration properties allowing cluster operators to alter that Redis configuration.

### How?
Follow [these instructions](http://bosh.io/docs/create-release.html) and refer to [this MySQL release](https://github.com/cloudfoundry/cf-mysql-release) as an example. If you get totally stuck, refer to this [community-created Redis release](https://github.com/cloudfoundry-community/redis-boshrelease).

### Expected Result
A BOSH release that facilitates the use of Redis by a CF deployed app.

### Resources
[http://bosh.io/docs/create-release.html](http://bosh.io/docs/create-release.html)