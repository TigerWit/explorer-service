# explorer-service

## 准备密钥文件
```
请参照DTN-docs说明，生成密钥文件certs
```
## 准备配置信息

### conf
```
进程运行监听的端口信息
httpport = 8080
```
### sdk.yaml.j2
```
根据配置模版sdk.yaml.j2
修改如下信息，得到sdk.yaml
client:
	organization:修改为自定义节点名称
tlsCerts:
	client:
		key:
			path:用户账户的私钥存放路径
		cert:
			path:用户账户的公钥证书存放路径
channels:
	tradechannel:
	peers:
		用户节点自定义名称
		endorsingPeer: 设置背书权限，默认false
        chaincodeQuery: 设置合约查询权限，默认false
        ledgerQuery: 设置账本查询权限，默认false
        eventSource: 设置事件权限，默认false
organizations:
  tiger.fx.com:
    mspid: 用户组织自定义mspid
    cryptoPath: 用户组织自定义用户账户mspid
    peers:
      - 用户节点自定义名称
peers:
  用户节点自定义名称:
    url: 用户节点自定义名称:监听端口，建议7051
    eventUrl: 用户节点自定义名称:监听端口，建议7053
    tlsCACerts: 用户节点的tls证书存放路径
```