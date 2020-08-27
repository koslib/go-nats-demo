# go-nats-demo

Me, playing around with [nats](https://nats.io/).

Just a simple playground for exploring nats, try proto code and experiment before actually using something in prod.

In the `main.go` file you can find examples for:

1. Setting up a consumer/producer for an async example
2. Setting up a request/reply model example for synchronous messages consumption


## Running it 
For running locally a nats server through docker: 
```
docker run --name nats -p 6222:6222 -p 4222:4222 -p 8222:8222 nats:latest
```

For a demonstration of the whole functionality (just logs printing, nothing fancy), you can run:
```
docker-compose up
```

and remember to rebuild if you change the code while experimenting:
```
docker-compose up --build
```

and when you're done tear down:
```
docker-compose down
```

> Note: the docker-compose setup includes two nats services, to showcase a cluster example.


## Nats-operator

### Deploy dummy service on k8s

#### Deploy with custom helm chart 

```
helm upgrade --install test http://www.koslib.com/mycharts/servicetpl-0.5.3.tgz -f .chart/values.yaml --wait --set image.tag=<your image tag>
``` 

#### Deploy with simple kubernetes deployment

```bash
kubectl create deployment test --image=<your local image name and tag here>
```

### Deploying the nats operator and cluster 

nats-operator [link](https://github.com/nats-io/nats-operator). Will do the namespace-scope installation.

When the operator has been installed, create a nats cluster:
```bash
cat <<EOF | kubectl create -f -
apiVersion: nats.io/v1alpha2
kind: NatsCluster
metadata:
  name: example-nats-cluster
spec:
  size: 1
  version: "1.3.0"
EOF
```

Also, install [nats streaming operator](https://github.com/nats-io/nats-streaming-operator) and create a nats streaming cluster:

```bash
echo '
---
apiVersion: "streaming.nats.io/v1alpha1"
kind: "NatsStreamingCluster"
metadata:
  name: "example-stan"
spec:
  size: 3
  natsSvc: "example-nats-cluster"
  config: {}
' | kubectl apply -f -
```


## key takeaways

* Sync messages support custom limits for what kind of overhead a client will get (eg. limit the amount of messages or the size of them)
* Async messages can define their binding functions
* Timeouts are configurable in pretty much anything
* Simple and clean Go APIs
* Wildcards support in topics/subjects
* Messages can be acked manually
* Message requeue 
* Message persistence supported through in-memory, file or sql stores (nats-streaming only)
* (k8s version) TLS support
* (k8s version) separate accounts and permissions support via k8s service accounts
* (k8s version) official helm chart for operator (good for prod use-cases)

## Disclaimer

This is a project only for demo-ing and showcase purposes. This is not an one-size-fits-all. Every use-case is 
different. I'd be happy to help if I can, just get in touch with me on [Twitter](https://www.twitter.com/koslib).
