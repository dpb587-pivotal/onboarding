---
title: Set up your CF target domain
labels:
- cf-deployment
type: story-feature
weight: 15
---

Set up your CF target domain
### What?
We're almost there. The final step is to target your newly deployed Cloud Foundry.

### How?
Run `cf api api.YOUR_DOMAIN.com --skip-ssl-validation`

### Expected Result
`cf login` works (your username is 'admin' and your password is in your generated `cf-deployment-vars.yml` as 'cf_admin_password')

### Troubleshooting
* Do the domain names associated with your TLS cert, your load balancer, and your `cf-deployment-vars.yml` system_domain all match the domain you registered on Freenom?
* Are the nameservers associated with your domain on Freenom the same as those associated with your load balancer? (run `bbl lbs` to check)
* When you `dig` or `ping` your CC api endpoint, is the IP you hit what you expect it to be?
* If you were redeploying, did you run the `bosh -n interpolate` step before running `bosh -d cf deploy`?

### Resources
[Forum question: Get Common Name from TLS cert](https://unix.stackexchange.com/questions/103461/get-common-name-cn-from-ssl-certificate)