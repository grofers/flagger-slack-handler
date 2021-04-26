
# flagger-slack-handler

Manual gating for Flagger based Canary deployments directly from Slack!

## Requirements

- [Flagger](https://flagger.app/)
- [flagger-loadtester](https://docs.flagger.app/usage/webhooks#load-testing)
- A Slack workspace




## Installation 

This project is intended to be run on Kubernetes. 

Install the `Deployment`:

```bash 
$ kubectl apply -n [NAMESPACE] -f https://raw.githubusercontent.com/grofers/flagger-slack-handler/master/kubernetes/deployment.yaml 
```

This `Deployment` needs to be exposed locally for Slack to interact with. If your cluster has capabilities for provisioning a public `LoadBalancer`, edit the below service to `LoadBalancer`. Alternatively, you may use expose the service using an Ingress controller.

```bash
$ kubectl apply -n [NAMESPACE] -f https://raw.githubusercontent.com/grofers/flagger-slack-handler/master/kubernetes/service.yaml 
```


    
## Run Locally

Follow the steps below to run this project locally

```bash
$ https://github.com/grofers/flagger-slack-handler
```

Go to the project directory

```bash
$ cd flagger-slack-handler
```

Build the binary

```bash
$ make build
```

Start the server

```bash
$ ./bin/flagger-slack-handler \
    -port=8080 \
    -loadtester-namespace=<namespace where flagger-loadtester is installed>
```

  
## Usage

### Pre-requisites

Before proceeding, install a [Slack command bot](https://api.slack.com/interactivity/slash-commands) in your Slack workspace and update the Request URL to point to the `LoadBalancer` / Ingress URL of your installation of `flagger-slack-handler`, suffixed with a `/handler`. 

Verify if the bot is running by issuing the following command from any Slack channel:

```
/flagger help
```

### Promotion

Run the following command to promote a Canary rollout:

```
/flagger promote [CANARY NAME] [CANARY NAMESPACE]

Example:

/flagger promote podinfo test
```

A promotion calls the `confirm-promotion` or `confirm-traffic-rollout` webhook. More info [here](https://docs.flagger.app/usage/webhooks).

### Rollback

Run the following command to rollback a Canary rollout:

```
/flagger rollback [CANARY NAME] [CANARY NAMESPACE]

Example:

/flagger rollback podinfo test
```

A rollback calls the `rollback` webhook. More info [here](https://docs.flagger.app/usage/webhooks).


## License

[MIT](https://choosealicense.com/licenses/mit/)

  