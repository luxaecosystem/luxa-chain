# LUXA Network (`luxa-1`)

Reference information for the LUXA Network sovereign Layer 1 blockchain.
Live network values in this document were checked on 2026-10-20.

## Network Specifications

- **Chain ID:** `luxa-1`
- **Source release tag:** [`v1.0.0`](https://github.com/luxaecosystem/luxa-chain/tree/v1.0.0)
- **Moniker:** `Raqeem`
- **Bech32 account prefix:** `luxa`
- **Coin type (SLIP-0044):** `118`
- **Base denomination:** `uluxa` (1 LUXA = 1,000,000 `uluxa`)
- **Secondary asset:** `SHARD`, base denomination `ushard` (1 SHARD = 1,000,000 `ushard`)
- **Consensus:** CometBFT v0.38.12
- **Cosmos SDK:** v0.50.10
- **Wasm runtime dependency:** `wasmd` v0.50.0
- **Genesis time:** `2026-10-08T14:59:10Z`
- **Genesis configuration:** CosmWasm's `wasm` module is enabled.

The chain ID, moniker, consensus version, and genesis time were checked against the public node. SDK and Wasm dependency versions are from the `v1.0.0` source release.

## Public Endpoints

- **CometBFT RPC:** <https://rpc.luxaecosystem.xyz>
- **Cosmos REST API (LCD):** <https://api.luxaecosystem.xyz>
- **Ping.pub Explorer:** <https://explorer.luxaecosystem.xyz/luxa>
- **LUXA Explorer:** <https://luxaecosystem.xyz/explorer.html>
- **Official website:** <https://luxaecosystem.xyz>

## Assets and Supply

The [public bank supply endpoint](https://api.luxaecosystem.xyz/cosmos/bank/v1beta1/supply?pagination.limit=1000) reported the following values when checked on 2026-10-20 (the latest block was around height `17241`):

| Asset | Base denom | On-chain supply | Display supply |
|---|---|---:|---:|
| LUXA | `uluxa` | `1002500000000000 uluxa` | 1,002,500,000 LUXA |
| SHARD | `ushard` | `10006000000000000000 ushard` | 10,006,000,000,000 SHARD |

These are observed supply values, not proof of an immutable maximum supply. The LUXA amount matches the often-quoted 1,002,500,000 LUXA figure, but this document does not claim that amount is a protocol-enforced hard cap.

The project documents an application-level conversion/forging rate of **10,000 SHARD = 1 LUXA**. This should not be interpreted as a native consensus-level exchange rate or a guaranteed on-chain swap.

## Governance

The Raqeem Council and ministerial reserve are described in project materials. On-chain multisig addresses and their relationship to any reserve should be independently verified before treating them as confirmed protocol properties.

## P2P Peer

The following peer ID and address were checked against the public RPC node. The P2P port was reachable during verification; peer availability is not guaranteed.

Add to `~/.luxa/config/config.toml`:

```toml
persistent_peers = "c6eb98f7a65a5fd23ef5c9004d792ec9a8e0b803@51.170.130.74:26656"
```

## Build from Source

### Prerequisites

- Go 1.25.0 (as declared by the `v1.0.0` source release)
- Make
- GCC or another supported C compiler

### Build and install

```bash
git clone https://github.com/luxaecosystem/luxa-chain.git
cd luxa-chain
git checkout v1.0.0
make install
```

Verify the installed binary:

```bash
luxad version --long
```

## Wallets

- **Keplr:** LUXA Network configuration is available in the [Keplr chain registry](https://github.com/chainapsis/keplr-chain-registry/blob/main/cosmos/luxa.json).
- **LUXA Vault ("Spider's Nest"):** [Chrome Web Store listing](https://chromewebstore.google.com/detail/luxa-vault-%E2%80%94-spiders-nest/jfhaclclifggggclkpeakpgddohpdbmm).
- **Leap:** Cosmos wallet support may require adding the network manually; check Leap's current network support before use.
