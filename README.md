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

## Running Prebuilt Binaries

1. Install [Barreto-Naehrig curves](https://github.com/dfinity/bn)

   1. Install the prebuilt libraries in their readme page into `/usr/lib`.

   1. Install the dependency libgmp `apt get install libgmp-dev`.

1. Download and run the prebuilt binaries (built on Ubuntu 16.04) from the [release page](https://github.com/helinwang/dex/releases).

## Build

### Build with Docker

```
$ docker-compose up -d
$ docker-compose exec dex bash
$ glide install
$ go test ./pkg/...
$ go build ./cmd/node/
```

### Build from Source

- Install the latest version of [Go](https://golang.org/doc/install#install)

- Install [Barreto-Naehrig curves](https://github.com/dfinity/bn)

  - Ubuntu or OSX can use the latest prebuilt libraries in the readme
    page.

  - Install the include files and built libraries into `/usr/include`
    and `/usr/lib` respectively (or anywhere else the Go build
    toolchain can find).

  - Install dependencies `apt install llvm g++ libgmp-dev libssl-dev`,
    they are required by cgo when building the BLS Go wrapper.

  - Test the installation by:
    ```
    $ go get github.com/dfinity/go-dfinity-crypto
    $ cd $GOPATH/src/github.com/dfinity/go-dfinity-crypto/bls
    $ go test
    ```

- Install package manager [Glide](https://glide.sh/)

- Download source and build
  ```
  $ go get github.com/helinwang/dex
  $ cd $GOPATH/src/github.com/helinwang/dex
  $ glide install
  $ go test ./pkg/...
  $ go build ./cmd/node/
  ```

## License

GPLv3

[1] Benchmark performed by running multiple nodes on my local machine, steps [here](./commands.md#pressure-testing). Machine configuration: 16core, 32GB. Please note these are preliminary results, the system has a lot of room for optimization. The block time can be more stable with some improvements.

[2] Normal operation is a likely event that happens when there is only one notarized block produced in the round. For more detail please see the [Dfinity Consensus Paper](https://dfinity.org/pdf-viewer/library/dfinity-consensus.pdf).
