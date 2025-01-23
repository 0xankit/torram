# Torram

Create a prototype for staking Runes for a basic Cosmos SDK chain called - Torram.

The staked asset must track its state and data availability between the Bitcoin network and Cosmos SDK chain - Torram at all times.

Example: Once the Rune (e.g. Liquidium) is staked in the Cosmos SDK chain - Torram, this transactional information should be visible in the Bitcoin network and Cosmos SDK chain - Torram. When it is unstaked, this transactional information should reflect in the Bitcoin network and Cosmos SDK chain - Torram.


## If assets on BTC side
1. When Rune should be staked on BTC it should be reflected on cosmos side
2. When Unstaked from BTC it should be reflected on cosmos-sdk side.


## Bitcoin Core
bitcoind --regtest
bitcoin.conf
```
rpcuser=yourusername
rpcpassword=yourpassword
regtest=1
txindex=1
server=1
```
bitcoin-cli --regtest --rpcuser=yourusername --rpcpassword=yourpassword createwallet "testwallet"

bitcoin-cli --regtest --rpcuser=yourusername --rpcpassword=yourpassword loadwallet "testwallet"
bitcoin-cli --regtest --rpcuser=yourusername --rpcpassword=yourpassword loadwallet "testwallet2"


bitcoin-cli --regtest --rpcuser=yourusername --rpcpassword=yourpassword listwallets

bitcoin-cli --regtest --rpcuser=yourusername --rpcpassword=yourpassword -rpcwallet=testwallet sendtoaddress "bcrt1qygqsu7qymllg2tc02rafzc06rymsh6h88srj94" 0.01

username --rpcpassword=yourpassword generatetoaddress 1 "$(bitcoin-cli --regtest --rpcuser=yourusername --rpcpassword=yourpassword -rpcwallet=testwallet2 getnewaddress )"

### List Unspent UTXO
bitcoin-cli --regtest --rpcuser=yourusername --rpcpassword=yourpassword -rpcwallet=testwallet2 listunspent


### Assumtions
1. Minimum unbonding time 50 blocks after this unbonding timelock expires.
2. Bitcoin staking workflow
```md
### From a Bitcoin staker's perspective, the Bitcoin staking protocol works as follows:

Staking bitcoin: the staker initiates the process by sending a staking transaction to the Bitcoin chain, locking her bitcoin in a self-custodian vault. More specifically, it creates a UTXO with two spending conditions:
1) timelock after which the staker can use her secret key to withdraw, and
2) burning this UTXO through a special extractable one-time signature (EOTS). In case of delegation, this EOTS belongs to the validator the stake delegates to.
Validation on PoS Chain: Once the staking transaction is confirmed on the Bitcoin chain, the staker (or the validator the staker delegates to) can start validating the PoS chain and signing votes valid blocks using the EOTS secret key. During her validation duty, there are two possible paths:
Happy Path: In the honest scenario, the staker follows the protocol and earns yield. The staker can then unbond via two approaches:
     1) wait for the existing timeclock to expire and then withdraw; or 2) submit an unbonding transaction to Bitcoin, which will unlock the bitcoin and return it to her after a parameterized unbonding period.
Unhappy Path: If the staker behaves maliciously, e.g., participates in double-spending attacks on the PoS chain, the staking protocol ensures her EOTS secret key is exposed to the public. Consequently, anyone can impersonate the staker to submit a slashing transaction on the Bitcoin chain and burn her bitcoin. This unhappy path ensures that safety violations are penalized, maintaining the overall integrity of the system.
```


### TO Stake
```
torram tx btcstaking stake "0de9390a2da201181d3f4250eedfc88574f074c78672382dcee4d44c2f5222fe" 1 --from user
```


Implementation Steps
Step 1: Track UTXO State in Torram
Add logic to Torram's staking module to maintain a mapping of staked UTXOs:

Key: txID:vout.
Value: {amount, cosmosAddress, isStaked, timestamp}.
Step 2: Use OP_RETURN Scripts on Bitcoin
When a BTC UTXO is staked:
Create a Bitcoin transaction with an OP_RETURN output containing:
Torram transaction hash (txHash).
Cosmos account address.
Metadata about the staking event.
Step 3: Build a Bitcoin Relayer
Create a relayer service to monitor Bitcoin transactions:
Use bitcoin-cli to track UTXO states (listunspent and getrawtransaction).
Parse OP_RETURN outputs to extract Torram staking/unstaking metadata.
Submit this data to Torram via MsgStakeBtc or MsgUnstakeBtc.
Step 4: Build a Torram Relayer
Create a relayer service to monitor Torram staking/unstaking events:
Use torramd query endpoints (staked-utxos).
Generate a Bitcoin transaction to update the UTXO state on the Bitcoin network (via OP_RETURN).
Step 5: Cryptographic Proofs
Include Bitcoin transaction proofs in Torram:
Use SPV (Simplified Payment Verification) proofs to validate Bitcoin transactions within Torram.
Submit these proofs with staking or unstaking requests to ensure authenticity.
Step 6: Data Availability
Ensure both chains have queryable endpoints:
Torram: Add queries like staked-utxo and staked-utxos to fetch staked BTC states.
Bitcoin: Use existing Bitcoin RPC calls (getrawtransaction, listunspent) for querying UTXO states.



#### TORRUM CLI
```bash
❯ torramd tx btcstaking stake-btc b8fca4fa6e990245c97170ca9c5d3aa439e09ea3a3ef6feaa3a34d3bcbfac53b 1 --amount 1trm --from alice
❯ torramd tx btcstaking unstake-btc b8fca4fa6e990245c97170ca9c5d3aa439e09ea3a3ef6feaa3a34d3bcbfac53b 1  --from alice
❯ torramd query btcstaking get-staked-btc b8fca4fa6e990245c97170ca9c5d3aa439e09ea3a3ef6feaa3a34d3bcbfac53b 1
```