# DEX

DEX is a decentralized exchange implementation that focuses on
low-latency and high-throughput.

- Block time: 1s on normal load, ~2.7s on high load [1]
- Near instant finality: a transaction is finalized after three block confirmations under normal operation [2]
- Transaction per second: ~2500 [1]

DEX implements the [Dfinity](https://dfinity.org/) consensus protocol described in their [paper](https://dfinity.org/pdf-viewer/library/dfinity-consensus.pdf).
And it implements native on-chain matching engine.

## Features

- Trading tokens
- Issue new tokens
- Sending, receiving, freezing and burning tokens

Please see [here](./commands.md) for how to run the nodes, use the wallet CLI, the
detailed steps of reproducing the features, and how to pressure test the system.

You can find the [White Paper](https://github.com/helinwang/dex/wiki/White-Paper) in the wiki. It has system overview and the plan to solve the scalability problem.

## Build

```bash
$ docker-compose up -d
$ docker-compose exec dex bash
$ glide install
$ go test ./pkg/...
$ go build ./cmd/node/
```

## API

```bash
$ # install echo framework
$ go get -u -f -v github.com/labstack/echo
$ go get -u golang.org/x/crypto/...
$ # run code
$ PORT=8080 go run api/*.go
$ # live reload
$ PORT=8080 make watch-api
```
