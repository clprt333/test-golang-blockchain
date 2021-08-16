To test:

Go to /tmp/ 
In 3 seperate command lines, assign 3 seperate nodeid's (env variable)
set NODEID=3000
set NODEID=4000
set NODEID=5000
copy blocks_3000
paste, rename to blocks_4000
paste, rename to blocks_5000
go run main.go createwallet (on all 3 nodes)
go run main.go createblockchain -address ADDRESS (on all 3 nodes)
go run main.go -from ADDRESS -to ADDRESS -amount 5 -mine (do this a couple times on all nodes) (send from address of 1 node to another)
now close 1 node, and reopen the same node with startnode -miner ADDRESS
You should see each of the nodes start mining new blocks.

Usage:
getbalance -address ADDRESS - get the balance of an address
createblockchain -address ADDRESS creates a blockchain and sends genesis reward to the address
printchain - Prints the blocks in the chain
send -from FROM -to TO -amount AMOUNT -mine - Sends a specified amount of coins. The -mine flag is set to mine from this node
createwallet - Creates a new Wallet
listaddresses - Lists the addresses in the wallet file
reindexutxo - Rebuilds the UTXO set
startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. -miner enables mining
