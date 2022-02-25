# explorer-service

## Prepare cryptographic files
```
Please refer to the DTN-docs instructions to generate the cryptographic files:certs.
```
## Prepare configuration information

### Configure the application
```
Configuration file：app.conf
You should customize this configuration file, although the sample configuration is available.
```
### Configure your blockchain sdk client
```
According to the configuration template:sdk.yaml.j2
Modify the following information to get sdk.yaml
client:
	organization:Custom Org Name
tlsCerts:
	client:
		key:
			path:Private key storage path for user
		cert:
			path:Public key certificate storage path for user 
channels:
	tradechannel:
	peers:
		Custom Peer Name
		endorsingPeer: Set up endorsement authority，default:false
        chaincodeQuery: Set up contract query authority，default:false
        ledgerQuery: Set up account query authority，default:false
        eventSource: Set event permissions，default:false
organizations:
  Custom Org Name:
    mspid: Custom Org mspid
    cryptoPath: Custom Org-User mspid
    peers:
      - Custom Peer Name
peers:
  Custom Peer Name:
    url: Custom Peer Name:Port，proposal 7051
    eventUrl: Custom Peer Name:Port，proposal 7053
    tlsCACerts: Custom Peer tlsca path
```
## Unzip the binary package
```
Select the appropriate compression package based on the computer computing architecture.
Here we provide a compact package of Linux and OSX platform architectures and verify the MD5 file:
.
├── explorer-service-darwin-amd64-1.0.tar.gz
├── explorer-service-darwin-amd64-1.0.tar.gz.md5
├── explorer-service-linux-amd64-1.0.tar.gz
└── explorer-service-linux-amd64-1.0.tar.gz.md5

```
### Unzip command can be like this
```
# In the explorer-service directory
cd $GOPATH/src/github.com/TigerWit/explorer-service
# Assuming you chose the Linux platform
tar zxvf explorer-service-linux-amd64-1.0.tar.gz
```
## Run explorer-service
```
cd $GOPATH/src/github.com/TigerWit/explorer-service
./explorer-service
```
## Wiki of the service APIs
```
If you succeed in starting explorer-service,There will be four API interfaces with HTTP forms to provide access support on your server machine.
```
### API:Query current state latest value based on key value
```
url: http://ip:port/querybykey
param: key
return: 
{
  status: {
    code: "status code"
    msg: "description information"
  },
  value: "transaction latest value"
}
```
### API:Querying historical information based on key value
```
url: http://ip:port/history
param: key
return: 
{
  status: {
    code: "status code"
    msg: "description information"
  },
  history: {
    {TxId:"transaction ID", Value:"transaction value", TimeStamp:"transaction timestamp"}
    ...
  }
}
```
### API:Query the transaction ID based on the key and value
```
url: http://ip:port/gettxidspec
param: key,value
return: 
{
  status: {
    code: "status code"
    msg: "description information"
  },
  tx_id: "transaction ID"
}
```
### API:Query transaction information based on transaction ID
```
url: http://ip:port/gettxbyid
param: txid
return: 
{
  status: {
    code: "status code"
    msg: "description information"
  },
  tx_info: {
    tx_id: "transaction ID",
    channel_id: "channel ID",
    block_info: {
      number: "block num",
      current_hash: "block hash",
      previous_hash: "previous block hash"
    },
    validation_code_name: "validation description",
    timestamp: {
      seconds: "timestamp seconds",
      nanos: "timestamp nanos"
    }
  }
}
```
