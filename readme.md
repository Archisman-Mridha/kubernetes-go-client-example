# Kubernetes Go Client Example

Using the [Kubernetes Go Client](https://github.com/kubernetes/client-go) library, we can interact with a Kubernetes cluster **through our go application**. For example, we can create or delete deployments, deploy CRDs and lots of other things. The primary idea is, the library will read a kube-config file and generate api clients for different api groups in the cluster.

This repository contains a simple example of how to create a deployment in your local Kubernetes cluster through a go application, using the Kubernetes go client library.

## Running the Application

After cloning this repository, run this command to download the dependency packages of our go workspace applications :

```sh
go work sync
```

Then use this command to run the `create deployment` application :

```sh
go run ./create-deployment
```

If you want to delete the created deployment, execute this command :

```sh
kubectl delete deployments demo-deployment
```