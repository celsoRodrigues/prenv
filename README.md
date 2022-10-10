# Self-Served Dynamic Environments for airflow

This project creates a webhook listener for pull requests and acts on every create and close of PR's

## Dependencies
- Docker
- kubernetes cluster

## Installation

* Create webhook Secret - This is the secret you will need to add to your github webhook configuration

```
make createsecret
```

* Deploy the watcher in your cluster

```
make install
```
