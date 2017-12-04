---
title: Use bbl to create a load-balancer
labels:
- cf-deployment
type: story-feature
weight: 9
---

Use bbl to create a load-balancer
### What?
Use `bbl` + your newly generated TLS certificate to create a **[load balancer](https://cloud.google.com/compute/docs/load-balancing/#types_of_load_balancing)** in front of your future Cloud Foundry deployment. CF doesn't require a load balancer, nor does it include one, but as you read the (P)CF docs you'll realize that it's assumed that you'll have one.

### How?
Run `bbl create-lbs` with the following parameters:

`--type`: Tells bbl which set of load balancers to create. Possible values are concourse and cf (use cf).
`--domain`: The domain that will resolve to your load balancer. It should match the domain that you registered on Freenom and the Common Name you entered when creating your TLS cert.
`--cert`: A path to a file with a PEM-encoded TLS certificate (the cert.pem you created in the last story). Remember that the cert should be valid for whichever domain you specify in the --domain flag.
`--key`: A path to a file with a PEM-encoded TLS key that corresponds to the cert provided in --cert.

All together that's: `bbl create-lbs --type cf --domain some-domain.com --cert cert.pem --key key.pem`

**NOTE:** If `bbl` fails to create your load balancer, run the command `bbl latest-error` to see what's up.

### Expected Result
Run `bbl lbs` to see load balancer details.

Visit Menu > Networking > Load balancing to see the load balancer that `bbl` generated on GCP.

Visit Menu > Networking > Cloud DNS to see the [Cloud DNS Zone](https://cloud.google.com/dns/overview) and [record sets](https://cloud.google.com/dns/overview#supported_dns_record_types) that `bbl` generated on GCP.

### Resources
[Docs: Types of load balancing](https://cloud.google.com/compute/docs/load-balancing/#types_of_load_balancing)
[Docs: What is cf-deployment?](https://github.com/cloudfoundry/cf-deployment/blob/master/deployment-guide.md)