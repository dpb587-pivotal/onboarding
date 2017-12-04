---
title: Connect to your BOSH environment
labels:
- bbl
type: story-feature
weight: 5
---

Connect to your BOSH environment
### What?
Time to target the Director VM you just `bbl up`ed.

### How?
To set the BOSH environment, run the following commands:

`bbl print-env` to see the environment variables you'll set to use your new BOSH. 

`eval "$(bbl print-env)"` to export them.

Run the following command to alias your BOSH environment, replacing `YOUR-ENV` with an environment name of your choice:

`bosh --environment="$BOSH_ENVIRONMENT" --ca-cert="$BOSH_CA_CERT" --client="$BOSH_CLIENT" --client-secret="$BOSH_CLIENT_SECRET" alias-env YOUR-ENV`

### Expected Results
When you run `bosh -e YOUR-ENV env`, it should print something like:

![image](https://www.pivotaltracker.com/file_attachments/77821681/download)

Annnd that's it!

# Congratulations, you've successfully deployed BOSH! #bigpatontheback