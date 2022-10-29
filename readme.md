### Scratch

1. Create a user account list to which stores user balances
9. *Implement the user account as wallets*
2. A feature to create and save the current state of chain
3. A feature to load a saved state of chain
4. Create multiple nodes to mine the blocks
5. Create a feature to send coins to other users
6. [API] Create a feature to view the balance of a user
7. Create a feature to view the transactions of a user
8. Create a feature to view the blocks 

10. Create a block explorer
11. Create a database to store all the accounts for reference which has balance>1 (to be used in the block explorer). It will also act as a node.
---

## Ideas
1. Use the NATS server to create a pub/sub system to push updates and syncronize the chain across all the nodes.
2. Create a central publisher for transactions to push them to mempool and again get status and push the transaction status to user, create a sql db to keep track of this transaction per user.

### Links

1. [Pos in Go](https://mycoralhealth.medium.com/code-your-own-proof-of-stake-blockchain-in-go-610cd99aa658)

### Notes

* Pos concept:
    How do we assign proper weights to them based on the number of tokens they staked?
    We fill our lotteryPool with copies of the validator’s address. They get a copy for each token they’ve staked. So a validator who put in 100 tokens will get 100 entries in the lotteryPool. A validator who only put in 1 token will only get 1 entry.
    We randomly pick the winner from our lotteryPool and assign their address to lotteryWinner.
* 1 block = max 15 transactions from top
* Websocket Message Types
| Type | Description |
| --- | --- |
| 1 | Client enter/exit |
| 2 | Transaction Published |