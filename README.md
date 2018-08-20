# explorer-service

## Prepare cryptographic files
```
Please refer to the DTN-docs instructions to generate the cryptographic files:certs.
```
## Prepare configuration information

### conf
```
Configuration file of port information for process operation monitoring：app.conf
eg.httpport = 8080
```
### sdk.yaml.j2
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
