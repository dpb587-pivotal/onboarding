---
title: Download and hook up your Concourse CLI
labels:
- concourse
type: story-feature
weight: 2
---

Download and hook up your Concourse CLI
### What?
Now that you have Concourse deployed, the first thing you'll want to do is download the `fly` CLI and authenticate with your target. This is done with the `fly login` command. The login command serves double duty: it authenticates with a given endpoint  and saves it under a more convenient name. The name and token are stored in ~/.flyrc (though you shouldn't really edit the file manually).

### How?
1. Download the `fly` CLI from the "no pipelines configured" page.
1. Move it to your computer's $PATH and make it executable.
1. Run `fly login --help` for instructions to login to Concourse.

### Expected Result
You successfully enter your credentials and the CLI prints "target saved."

### Resources
[Fly login documentation](https://concourse.ci/fly-login.html)

### Relevant Repos and Teams
[concourse/fly](https://github.com/concourse/fly)