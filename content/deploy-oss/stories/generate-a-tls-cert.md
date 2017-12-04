---
title: Generate a TLS cert
labels:
- cf-deployment
type: story-feature
weight: 8
---

Generate a TLS cert
### What?
In the next story, we're going to use `bbl` to create a load balancer for our soon-to-be Cloud Foundry deployment. This load balancer requires a TLS certificate, so we'll create one now.

### How?
We're going to use [OpenSSL](https://www.openssl.org/) to generate a self-signed TLS certificate. In real life, you'd get one from a trusted Certificate Authority like [Let's Encrypt](https://letsencrypt.org/), but for this exercise self-signed is sufficient.

To do this, run `openssl req` with a few arguments:

* `-x509`outputs a x509 structure, a standard that defines the format of public key certificates.
* `-newkey rsa:2048`generates a new RSA key of 2048 bits in size.
* `-keyout` and `-out` arguments set output file paths (I generally use `key.pem` and `cert.pem`, respectively)
* The `-nodes` argument sets it to not encrypt private keys, meaning you won't have to enter a PEM passphrase.

**Caution:** when you run OpenSSL, it will ask you to enter values like your Country and State. The only one that really matters (and it matters a great deal) is the "Common Name," which you should fill with the Cloud Foundry system domain you registered in the last story (e.g. your-domain.com).

### Expected Result
You should have two new `.pem` files, one containing your private key and one containing your cert.

When you run `openssl x509 -noout -subject -in cert.pem` the information returned matches what you entered.

### Resources
[Blog post: Signed vs Self-Signed Certificates](https://www.thoughtco.com/signed-vs-self-signed-certificates-3469534)
[Blog post: SSL vs TLS, what's the difference?](https://luxsci.com/blog/ssl-versus-tls-whats-the-difference.html)
[Tutorial: How SSL and TLS encryption works](http://computer.howstuffworks.com/encryption4.htm)
[How to: Get Common Name from TLS cert](https://unix.stackexchange.com/questions/103461/get-common-name-cn-from-ssl-certificate)