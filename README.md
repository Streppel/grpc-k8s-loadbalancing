Quick and dirty experiment to see how gRPC clients interact with different types of kubernetes services. The main goal here is to analyze load balancing capabilities.

# How to run
```bash
make create-cluster && make load-images && make deploy
```

Depends on having [kind](https://kind.sigs.k8s.io/) installed.

Then inspect the cluster and observe the pods logs.
