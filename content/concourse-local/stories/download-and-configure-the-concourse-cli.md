---
title: Download and configure the Concourse CLI
labels:
- concourse
type: story-feature
weight: 3
---

Download and configure the Concourse CLI
### What?
Now that you have Concourse deployed, the first thing you'll want to do is download the `fly` CLI and authenticate with your target. This is done with the `fly login` command. The login command serves double duty: it authenticates with a given endpoint  and saves it under a more convenient name. The name and token are stored in ~/.flyrc (though you shouldn't really edit the file manually).

### How?
1. Download the `fly` CLI from the "no pipelines configured" page.
1. Make the program executable by running `chmod +x fly` on the file.
1. Move the executable to your computer's $PATH.

> The `$PATH` environment variable sets a selection of directories from which executable files can be run directly without requiring a full path (e.g. running `fly`, not `~/Downloads/fly`). If you `echo $PATH` you can see that variable. In all likelihood one of the directories is `/usr/local/bin/`. Unless you have another preference, move the `fly` executable to that directory.

To login to Concourse, run `fly -t lite login -c http://192.168.100.4:8080`

### Expected Result
You've successfully entered your credentials and the CLI prints "target saved."

Run `fly login --help` to learn about other commands.

### Resources
[Fly login documentation](https://concourse.ci/fly-login.html)
[Cheatsheet: how Unix file permissions work](https://danflood.com/tech-stuff/chmod-cheat-sheet/)

### Relevant Repos and Teams
[concourse/fly](https://github.com/concourse/fly)