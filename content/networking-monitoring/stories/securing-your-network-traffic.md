---
title: Securing your network traffic
labels:
- security
type: story-feature
weight: 2
---

Securing your network traffic
### What?
The `public_networks` (or `all_pcfdev`) security group allows broad access to your CF deployment's network.  A malicious user could likely access internal systems that would not normally be exposed to the public internet. Thus, in a production environment you wouldn't have such open security groups.

Let's verify this access and secure it.

### How?
1. `cf ssh` to an application's container.
1. Download the **[nmap](https://nmap.org/download.html)** tool (the .tar.bz2 is your best bet, as you can run it as an executable file without requiring root priviledges to install it).
1. Scan a section of the private network by running:
`./nmap -v YOUR-IP/24`
1.  Find a IP with an open HTTP port (80) and curl that IP address:
`curl -H "Host: dora.<YOUR-IP>.xip.io" http://<OPEN-IP>`
1. As an admin, unbind `public_networks` (or `all_pcfdev`) from the running security groups and restart your application.
1. Run the scan again and try to hit the same IP with the open HTTP port.

### Expected Result
The initial scan and curl should succeed. After removing the security group and restarting your application curling the open port should fail.

Poke around to see what else you can and can't do without that Security Group. Remember to re-bind it at the end or, if you'd like, try creating and testing your own based on those [most commonly used](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html#typical-groups).

### Resources
[Nmap Security Scanner](https://nmap.org/) 
[SecTools.Org: Top 125 Network Security Tools](http://sectools.org/)
[TCP and UDP Ports Explained](https://www.bleepingcomputer.com/tutorials/tcp-and-udp-ports-explained/)
[Typical Application Security Groups](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html#typical-groups)