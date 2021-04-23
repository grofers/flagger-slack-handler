# flagger-slack-handler

Manage manual gating for Flagger based Canary deployments directly from Slack!

## Installation

### Kubernetes

```bash
kubectl apply -n [YOUR NAMESPACE] -f kubernetes/
```

Make sure to edit the `loadtester-namespace` arugement in `kubernetes/deployment.yaml` to point to the namespace where your `flagger-loadtester` is installed.

### Locally

Before proceeding, make sure your current `kubeconfig` is set to point to the cluster where `flagger` and `flagger-loadtester` are installed.

```bash
$ CGO_ENABLED=0 go build -a -o flagger-slack ./cmd

$ ./flagger-slack \
    -port=8080 \
    -loadtester-namespace=flagger
```

