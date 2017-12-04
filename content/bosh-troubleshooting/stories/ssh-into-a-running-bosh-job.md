---
title: SSH into a running BOSH job
labels:
- bosh operator
type: story-feature
weight: 2
---

SSH into a running BOSH job
### What?
To SSH into a BOSH job, you need to use your BOSH Director as a gateway host (like a proxy).

### How?
Set up your gateway host by passing a few flags into your `bosh ssh` command or by setting environment variables.
* `--gw-user=` or $BOSH_GW_USER should be set to `jumpbox`.
* `--gw-host=` or $BOSH_GW_HOST should be set to your Bosh Director's IP address.
* `--gw-private-key=` or $BOSH_GW_PRIVATE_KEY is a little more complicated. You'll need to print your Bosh SSH key into a file by running `bbl ssh-key > bosh.pem` and using the path of that file as the environment variable or argument value (you may have to adjust the `.pem` file permissions for this to work).

To remember the required commands/environmental variable names you can run `bosh ssh --help`.

### Expected Result
Running `bosh -d cf ssh ...` opens a shell in your targetted machine.

### Resources
[Forum question: What's the distinction between an HTTP proxy, tunnel, and gateway?](http://stackoverflow.com/questions/10377679/whats-distinction-of-http-proxy-tunnel-gateway)