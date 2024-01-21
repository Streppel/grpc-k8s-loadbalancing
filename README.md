Quick and dirty experiment to see how gRPC clients interact with different types of kubernetes services. The main goal here is to analyze load balancing capabilities.

Inspired by [this](https://github.com/jtattermusch/grpc-loadbalancing-kubernetes-examples#example-1-round-robin-loadbalancing-with-grpcs-built-in-loadbalancing-policy).

# How to run
```bash
make create-cluster && make load-images && make deploy
```

Depends on having [kind](https://kind.sigs.k8s.io/) installed.

Then inspect the cluster and observe the pods logs.
