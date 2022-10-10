# Self-Served Dynamic Environments for airflow

This project creates a webhook listener for pull requests and acts on every create and close of PR's

## Dependencies
- Docker
- kubernetes cluster

## Installation
* Configure gihub webhook

- Under your repo webhook section, create a webhook pointing to your server's ingress
- Add your hostname in the values.yaml file.

* Create webhook Secret - This is the secret you will need to add to your github webhook configuration

```
export HOOK="your_github_webhook_password"
make createsecret
```

* Deploy the watcher in your cluster

```
make install
```
