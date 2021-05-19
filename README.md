# go-aoa


### what's the aurora

The Aurorachain project has three phases, this is the first phase of the code.

### Get the code
This repo only can clone by SSH,For example:
```
git clone git@github.com:aurorachain-io/go-aoa.git
```

### Building the source

For prerequisites and detailed build instructions please read the [Installation Instructions](https://github.com/ethereum/go-ethereum/wiki/Building-Ethereum) on the wiki.

Building `aoa` requires both a Go (version 1.13 or later) and a C compiler. You can install
thaoa using your favourite package manager. Once the dependencies are installed, run

```shell
make aoa
```

or, to build the full suite of utilities:

```shell
make all
```

### Deploy

Under Linux or Mac, get the compressed file from release and extract it to get the executable file aoa.Create the storage directory /data/aoa, and copy the executable file to the directory. Then execute the startup command, in which '--port' is the customized chain synchronization port, '--rpc' would open RPC service, '--rpcaddr' is the self-defined RPC listening IP which should set to 127.0.0.1 if you do not want the raoaote RPC connection, --rpcport is the customized RCP listening port.
for example
```
tar zxvf aoa-linux-amd64-1.1.16-unstable.tar.gz
mkdir –p /data/aoa
cp aoa-linux-amd64-1.1.16-unstable/aoa /data/aoa/
nohup /data/aoa/aoa --datadir /data/aoa/aoa-data --port 30303 --rpc --rpcaddr 0.0.0.0 --rpcport 8545 2>> /data/aoa/aoa.log &
```
### Attach the console
```
/data/aoa/aoa attach /data/aoa/aoa-data/aoa.ipc
```
### Common console command
```
#get block height
aoa.blockNumber
#get block info
aoa.getBlock(blockHashOrBlockNumber)
#get accounts in wallet
aoa.accounts
#get transaction info
aoa.getTransaction(transactionHash)
#generate new accounts and store thaoa in the keystore directory, encrypted with passphrase
personal.newAccount(passphrase)
#sent transaction
personal.sendTransaction({from:'affress',to:'address',value:web3.toWei(100,'aoa'),action:0}, "password")
#start rpc by console
admin.startRPC("0.0.0.0", 8545)
#stop rpc by console
admin.stopRPC()
```
