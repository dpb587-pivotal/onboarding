---
title: Create an admin user in UAA using the uaac CLI
labels:
- security
type: story-feature
weight: 2
---

Create an admin user in UAA using the uaac CLI
### What?
The UAA is the identity management service for Cloud Foundry. Its primary role is as an OAuth2 provider, issuing tokens for client applications to use when they act on behalf of Cloud Foundry users. In collaboration with the login server, it can authenticate users with their Cloud Foundry credentials, and can act as an SSO service using those credentials (or others). It has endpoints for managing user accounts and for registering OAuth2 clients, as well as various other management functions.

Administrators generally create users using the cf CLI, which creates user records in the UAA _and_ associates them with org and space roles in the CCDB. However, there are occasions where you'd want to bypass Cloud Controller and only create a user in UAA. When that need arises, you can use the uaac CLI.

### How?
Following **[these instructions](https://docs.cloudfoundry.org/adminguide/uaa-user-management.html)**, create a new admin user in UAA using the uaac CLI.

### Expected Result
Potentially a new UAA user. Potentially an inability to use uaac due to an invalid token.

### Resources
[Docs: Creating and Managing Users with the UAA CLI (UAAC)](https://docs.cloudfoundry.org/adminguide/uaa-user-management.html)
[Docs: UAA SysAdmin Guide](https://github.com/cloudfoundry/uaa/blob/master/docs/Sysadmin-Guide.rst)
[Blog post: Introducing the UAA and Security for Cloud Foundry](https://blog.pivotal.io/pivotal-cloud-foundry/products/introducing-the-uaa-and-security-for-cloud-foundry) (written in '12)

### Relevant Repos and Teams
**UAA:** [cloudfoundry/cf-uaac](https://github.com/cloudfoundry/cf-uaac)