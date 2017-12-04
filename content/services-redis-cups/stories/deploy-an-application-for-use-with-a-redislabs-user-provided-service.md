---
title: Deploy an application for use with a RedisLabs user-provided service
labels:
- app-dev
- services
type: story-feature
weight: 2
---

Deploy an application for use with a RedisLabs user-provided service
### What?
To play around with the RedisLabs user-provided service instance we just created, we're going to need an app to use it. Let's deploy and bind one now!

### How?
Use this [sample app](https://github.com/pivotal-cf/cf-redis-example-app) and the [instructions](https://github.com/pivotal-cf/cf-redis-example-app#using-redis-labs-cups) provided in the `README` to deploy the sample application and bind your service instance to it.

### Expected Result
You can see the sample app listed under `cf apps`.
You can see the service listed under `cf services`, and see that it is bound to the sample app.

From your terminal, run:
```
export APP=APP_NAME.YOUR_CF_DOMAIN.com
curl -X PUT $APP/foo -d 'data=some-value'
```
This should return `success`. Next run:
```
curl -X GET $APP/foo
```
This should return `some-value`.

The value should also be present in your RedisLabs Dashboard.

### Resources
[User Provided Services (CUPS) Documentation](https://docs.cloudfoundry.org/devguide/services/user-provided.html)
[Bind a Service Instance Documentation](https://docs.cloudfoundry.org/devguide/services/application-binding.html)