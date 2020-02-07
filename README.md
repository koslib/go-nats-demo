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

## key takeaways

* Sync messages support custom limits for what kind of overhead a client will get (eg. limit the amount of messages or the size of them)
* Async messages can define their binding functions
* Timeouts are configurable in pretty much anything
* Simple and clean Go APIs

## Disclaimer

Of course this is not intended for production. Every use-case is different. I'd be happy to help if I can, just get
in touch with me on [Twitter](https://www.twitter.com/koslib).