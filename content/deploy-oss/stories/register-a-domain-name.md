---
title: Register a domain name
labels:
- cf-deployment
type: story-feature
weight: 7
---

Register a domain name
### What?
Now that you've done your victory dance, on to bigger and better things&mdash;like getting a domain name for your Cloud Foundry deployment.

We need this name before deployment because the load balancer we're about to create and the TLS cert for that load balancer both need to know the domain that they're serving. Also, BOSH needs to know this domain before deployment so that it can associate subdomains with each CF component. Cloud Controller, for instance, is `api.your-domain.com`.

### How?
Register a domain name for free with **[Freenom](http://www.freenom.com/en/index.html?lang=en)**. We do not know the nameservers for your load balancer because we haven't created it yet, so hang tight on setting up the DNS.

### Expected Result
Freenom says that you have successfully registered your domain. Should look something like this:

![Imgur](http://i.imgur.com/E3WnV24.png)

### Resources
[Blog post: What is a domain name?](https://www.lifewire.com/what-is-a-domain-name-2483189)