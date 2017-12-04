---
title: Start PCF Dev and login
labels:
- core tools
type: story-feature
weight: 4
---

Start PCF Dev and login
### What?
You've already downloaded PCF Dev from Pivotal Network. Now you're going to run it!

### How?
Unzip the file you downloaded in the first story (e.g. `unzip pcfdev-VERSION-linux.zip`)

Run the executable binary file inside (e.g. `./pcfdev-VERSION-linux`)

From the command line, run `cf dev start`. Because you are running it for the first time it will download the image and import it to VirtualBox before starting the VM. In the future, the same command will start PCF Dev without downloading or importing it again.

FYI ...this process takes awhile. How long of a while will be influenced by your internet connectivity and bandwidth. Go get a snack or read some of the links.

When it wraps up, PCF Dev will have printed credentials for two users, `user` and `admin` and the command you should use to login. The users have different permissions, so make sure to choose `admin`, then select the `pivotal` org when prompted.

### Expected Result
Run `cf target`. You'll see a line that says `API endpoint:   https://api.local.pcfdev.io`, followed by your user, org, and space info.

### Resources
[Tutorial: Getting Started with PCF Dev](https://pivotal.io/platform/pcf-tutorials/getting-started-with-pivotal-cloud-foundry-dev/introduction)
[Blog post: Meet Pivotal Cloud Foundry Dev, your Ticket To Running Cloud Foundry Locally](https://content.pivotal.io/blog/meet-pcf-dev-your-ticket-to-running-cloud-foundry-locally)
[Blog post: A little diddy about binary file formats](https://betterexplained.com/articles/a-little-diddy-about-binary-file-formats/)

### Relevant Repos and Teams
**PCF Dev:** [pivotal-cf/pcfdev](https://github.com/pivotal-cf/pcfdev),
**Also PCF Dev:** [pivotal-cf/pcfdev-cli](https://github.com/pivotal-cf/pcfdev-cli)
