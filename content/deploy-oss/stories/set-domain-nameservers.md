---
title: Set domain nameservers
labels:
- cf-deployment
type: story-feature
weight: 10
---

Set domain nameservers
### What
A nameserver is a server on the internet that specializes in handling queries regarding the location (e.g. IP) of a domain name’s content. The domain you registered on Freenom needs to point at the nameservers associated with your load balancer on GCP.

### How
Go to your **[Freenom dashboard](https://my.freenom.com/clientarea.php)**, select Services > My Domains, then click the Manage Domain button. From the Management Tools dropdown, select Nameservers. Fill the text boxes with the nameservers associated with your load balancer (you can find these by running `bbl lbs`). Remember to select the "Use custom nameservers (enter below)" radio button!

### Expected Result
Go to your GCP Cloud DNS entry. Find the IP associated with `*.your-domain.com`. Running `dig api.your-domain.com` should return the same IP address.

Remember, propagation of the DNS changes may take a little while, so it might not work immediately.

### Resources
[Blog post: What is a Domain Name Server?](https://www.lifewire.com/what-is-a-dns-server-817513) 
[Video: How DNS works](https://www.youtube.com/watch?v=GlZC4Jwf3xQ)