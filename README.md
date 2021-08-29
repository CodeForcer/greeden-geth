# Greeden-Geth

Greeden-Geth is a protocol-agnostic client which uses a greedy algorithm to pick the most profitable blocks to submit to the network out of Flashbots, Eden, and regular mining.

Miners running Eden-Geth are losing large amounts of profits by publishing Flashbots blocks tainted with Eden slots, reducing their MEV rewards. These losses often total more than 10 ETH per block.

This client removes the lack of transparency, allowing miners to obtain increased rewards from their work.

PR's are welcome, and we invite both the Flashbots and Eden community to help us improve this client. This client is protocol agnostic, and my only aim is maximise miner profitability.

## Quick start

```
git clone https://github.com/CodeForcer/greeden-geth
cd greeden-geth
make geth
```

See [here](https://geth.ethereum.org/docs/install-and-build/installing-geth#build-go-ethereum-from-source-code) for further info on building Greeden-Geth from source.
