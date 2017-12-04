---
title: Install the Bosh v2 CLI
labels:
- concourse
type: story-feature
weight: 1
---

Install the Bosh v2 CLI
### What?
In this next section you are going to deploy Concourse locally in a single VM using `bosh create-env`, for which you'll need to install the Bosh v2 CLI.

### How?
Install the [Bosh v2 CLI](https://bosh.io/docs/cli-v2.html#install) with brew by running `brew install bosh-cli`. If you find yourself having trouble with the brew package manager, [these are the instructions to install manually](https://bosh.io/docs/cli-v2.html#install).

### Expected Result
Running `bosh -v` should print "version 2.0.33-...." or above.