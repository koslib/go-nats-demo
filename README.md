# go-nats-demo

Me, playing around with [nats](https://nats.io/).

Just a simple playground for exploring nats, try proto code and experiment before actually using something in prod.

For running locally a nats server through docker: 
```
docker run --name nats -p 6222:6222 -p 4222:4222 -p 8222:8222 nats:latest
```

## key takeaways

* Sync messages support custom limits for what kind of overhead a client will get (eg. limit the amount of messages or the size of them)
* Async messages can define their binding functions
* Timeouts are configurable in pretty much anything
* Simple and clean Go APIs