# Eden-Geth

Eden is an optional, non-consensus breaking transaction ordering protocol for Ethereum blocks which allows network participants to guarantee placement and protection from arbitrary reordering. The system offers a transparent and fair set of rules to order transactions within each block. An accompanying token reward system realizes MEV profits to block producers to maximize network security.

## Quick start

```
git clone https://github.com/eden-network/eden-geth
cd eden-geth
make geth
```

See [here](https://geth.ethereum.org/docs/install-and-build/installing-geth#build-go-ethereum-from-source-code) for further info on building Eden-Geth from source.

## Documentation

See [here](https://docs.flashbots.net) for Flashbots documentation.

| Version | Spec                                                                                        |
| ------- | ------------------------------------------------------------------------------------------- |
| v0.3    | [MEV-Geth Spec v0.3](https://docs.flashbots.net/flashbots-auction/miners/mev-geth-spec/v03) |
| v0.2    | [MEV-Geth Spec v0.2](https://docs.flashbots.net/flashbots-auction/miners/mev-geth-spec/v02) |
| v0.1    | [MEV-Geth Spec v0.1](https://docs.flashbots.net/flashbots-auction/miners/mev-geth-spec/v01) |

### Feature requests and bug reports

If you are a user of MEV-Geth and have suggestions on how to make integration with your current setup easier, or would like to submit a bug report, we encourage you to open an issue in this repository with the `enhancement` or `bug` labels respectively. If you need help getting started, please ask in the dedicated [#⛏️miners](https://discord.gg/rcgADN9qFX) channel in our Discord.
See [here](https://docs.edennetwork.io) for Eden Network documentation. Block producers can [join](https://docs.edennetwork.io/for-block-producers/getting-started) Eden Network in three easy steps (and running `eden-geth` is step number one!)
