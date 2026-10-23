# LUXA Network (`luxa-1`)

Reference information for the LUXA Network sovereign Layer 1 blockchain.
Live network values in this document were checked on 2026-10-23 against the public node and endpoints (latest block around height `22558`). Each section states what it covers.

## Network Specifications

- **Chain ID:** `luxa-1`
- **Source release tag:** [`v1.0.0`](https://github.com/luxaecosystem/luxa-chain/tree/v1.0.0)
- **Moniker (public node):** `Raqeem`
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

The [public bank supply endpoint](https://api.luxaecosystem.xyz/cosmos/bank/v1beta1/supply?pagination.limit=1000) reported the following values when checked on 2026-10-23 (the latest block was around height `22558`):

| Asset | Base denom | On-chain supply               | Display supply           |
| ----- | ---------- | ----------------------------- | ------------------------ |
| LUXA  | `uluxa`    | `1002500000000000 uluxa`      | 1,002,500,000 LUXA       |
| SHARD | `ushard`   | `10006000000000000000 ushard` | 10,006,000,000,000 SHARD |

These are observed supply values, not proof of an immutable maximum supply. The LUXA amount matches the often-quoted 1,002,500,000 LUXA figure, but this document does not claim that amount is a protocol-enforced hard cap.

The project documents an application-level conversion/forging rate of **10,000 SHARD = 1 LUXA**. This should not be interpreted as a native consensus-level exchange rate or a guaranteed on-chain swap.

## Staking Parameters

Values read from the node on 2026-10-23 (latest block around height `22558`):

| Parameter               | Value   |
| ----------------------- | ------- |
| Inflation               | 0%      |
| Unbonding time          | 21 days |
| Double-sign slashing    | 5%      |
| Downtime slashing       | 1%      |
| Downtime window         | 100 blocks, minimum 50% signed |
| Downtime jail duration  | 10 minutes |
| Maximum validators      | 100     |

## Validators

Active validator set on 2026-10-23 (2 validators):

| Rank | Moniker | Voting power     | Share | Commission | Website                        |
| ---- | ------- | ---------------- | ----- | ---------- | ------------------------------ |
| 1    | Raqeem  | 3,100,001 LUXA   | 75.6% | 10%        | <https://luxaecosystem.xyz>    |
| 2    | Luqman  | 1,000,504 LUXA   | 24.4% | 10%        | <https://luxaecosystem.xyz>    |

### Raqeem

- **Account address:** `luxa1qgx524exntta5r22knd9e78k5vyehhl46xprty`
- **Operator address:** `luxavaloper1qgx524exntta5r22knd9e78k5vyehhl4j2vskq`
- **Hex address:** `BE89691113B143F695D7AC35FE88086D7AAD8522`
- **Consensus address:** `luxavalcons1h6ykjygnk9pld9wh4s6lazqgd4a2mpfz5jqx6f`
- **Consensus public key:** `{"@type":"/cosmos.crypto.ed25519.PubKey","key":"Rje/KiQRcFVyTJBDayloPFP/LjVgnHje2GbArmlM3z4="}`

### Luqman

- **Account address:** `luxa13h4u8hrc5v3md6q20xultre44vqtr822cxs88z`
- **Operator address:** `luxavaloper13h4u8hrc5v3md6q20xultre44vqtr822s2a56x`
- **Hex address:** `03BF5910C09FA81CBA3E6EC67D005EF8CB0DB1C9`
- **Consensus address:** `luxavalcons1qwl4jyxqn75pew37dmr86qz7lr9smvwf0lulxw`
- **Consensus public key:** `{"@type":"/cosmos.crypto.ed25519.PubKey","key":"mc0kw77PZq2et0iKu6sqkn7s9OO0F1AQKBwBoPHXzyY="}`

Voting power is concentrated in two validators. Independent verification of the validator set can be done with `luxad q staking validators` against the public endpoints.

## IBC Bridges

The ICS-20 (`transfer`) channels below are reported as `STATE_OPEN` by the node (checked on 2026-10-23).

| Parameter           | Osmosis                       | Cosmos Hub                    |
| ------------------- | ----------------------------- | ----------------------------- |
| Counterparty chain  | `osmosis-1`                   | `cosmoshub-4`                 |
| LUXA channel        | `channel-1`                   | `channel-2`                   |
| Counterparty channel| `channel-111100`              | `channel-1941`                |
| Port                | `transfer`                    | `transfer`                    |
| IBC path of `uluxa` | `transfer/channel-111100/uluxa` | `transfer/channel-1941/uluxa` |
| Test transfer block | #73,021,145                   | #33,538,905                   |

### LUXA <-> Osmosis

- **Connection:** `connection-1` (LUXA) <-> `connection-11099` (Osmosis)
- **Tendermint client:** `07-tendermint-4` (LUXA) <-> `07-tendermint-3760` (Osmosis)
- **LUXA denom on Osmosis:** `ibc/8965A5F178427372659F3DD28D1F40770A905038A48DFB5D9C3800A5D81BE5DD`
- **Test transfer:** [`F3228216EA18C2599ADB0BFFEF37B886FF1D469B6669280BFD986CDD23C625A0`](https://www.mintscan.io/osmosis/tx/F3228216EA18C2599ADB0BFFEF37B886FF1D469B6669280BFD986CDD23C625A0)

### LUXA <-> Cosmos Hub

- **Connection:** `connection-2` (LUXA) <-> `connection-1313` (Cosmos Hub)
- **Tendermint client:** `07-tendermint-5` (LUXA) <-> `07-tendermint-1519` (Cosmos Hub)
- **LUXA denom on Cosmos Hub:** `ibc/36D60CB051F8068F30CB57A356C5FF15ABAF9A25DD2AD91BCD40A81056F05E8A`
- **Test transfer:** [`596D09329815E09A69BA851E061CD6F7AFB69BD5F0E1C9E2835C87A20644F78B`](https://www.mintscan.io/cosmos/tx/596D09329815E09A69BA851E061CD6F7AFB69BD5F0E1C9E2835C87A20644F78B)
- **Asset page:** [Mintscan](https://www.mintscan.io/cosmos/assets/ibc/36D60CB051F8068F30CB57A356C5FF15ABAF9A25DD2AD91BCD40A81056F05E8A)
- **Relayer signer on Cosmos Hub:** `cosmos1n2udk7p3qxeqywrg7nk5ldq28dwjclh0h532te`

### Manual IBC transfers (Keplr, Advanced IBC Transfer)

| Direction                  | Destination chain | Source channel (from LUXA) | Source channel (return path) |
| -------------------------- | ----------------- | -------------------------- | ---------------------------- |
| LUXA -> Osmosis            | Osmosis           | `channel-1`                | `channel-111100`             |
| LUXA -> Cosmos Hub         | Cosmos Hub        | `channel-2`                | `channel-1941`               |

IBC transfers rely on a relayer being online. Transfers can stay pending if no relayer is running, so check the relayer status before sending a large amount. Values of an IBC-wrapped LUXA on another chain depend on that chain's own markets and are not guaranteed by this document.

## Governance

The Raqeem Council and ministerial reserve are described in project materials. On-chain multisig addresses and their relationship to any reserve should be independently verified before treating them as confirmed protocol properties.

## P2P Peer

The following peer ID and address were checked against the public RPC node. The P2P port was reachable during verification; peer availability is not guaranteed.

Add to `~/.luxa/config/config.toml`:

```
persistent_peers = "c6eb98f7a65a5fd23ef5c9004d792ec9a8e0b803@51.170.130.74:26656"
```

## Build from Source

### Prerequisites

- Go 1.25.0 (as declared by the `v1.0.0` source release)
- Make
- GCC or another supported C compiler

### Build and install

```
git clone https://github.com/luxaecosystem/luxa-chain.git
cd luxa-chain
git checkout v1.0.0
make install
```

Verify the installed binary:

```
luxad version --long
```

## Wallets

- **Keplr:** LUXA Network configuration is available in the [Keplr chain registry](https://github.com/chainapsis/keplr-chain-registry/blob/main/cosmos/luxa.json).
- **LUXA Vault ("Spider's Nest"):** [Chrome Web Store listing](https://chromewebstore.google.com/detail/luxa-vault-%E2%80%94-spiders-nest/jfhaclclifggggclkpeakpgddohpdbmm).
- **Leap:** Cosmos wallet support may require adding the network manually; check Leap's current network support before use.
