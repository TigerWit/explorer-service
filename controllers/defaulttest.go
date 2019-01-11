package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	peer3 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/common/util"
	cb "github.com/hyperledger/fabric/protos/common"
)

func (c *MainController) SealTest() {
    key := c.GetString("key")
    value := c.GetString("value")
    data := make(map[string]string)
    sdk, err := fabsdk.New(config.FromFile("./sdk_test.yaml"))
    if err != nil {
        data["ERROR"] = fmt.Sprintf("Failed to create new SDK: %s", err)
        c.Data["json"] = data
        c.ServeJSON()
        return
    }
    defer sdk.Close()

    channelProvider := sdk.ChannelContext("tradechanneltest",
        fabsdk.WithUser("Admin"),
        fabsdk.WithOrg("tiger.fx.com"))

    channelClient, err := channel.New(channelProvider)
    if err != nil {
        data["ERROR"] = fmt.Sprintf("create channel client fail: %s", err)
        c.Data["json"] = data
        c.ServeJSON()
        return
    }

    args := [][]byte{[]byte(key), []byte(value)}
    _, err = channelClient.Execute(channel.Request{ChaincodeID: "sealtx", Fcn: "seal", Args: args},
        channel.WithRetry(retry.DefaultChannelOpts))
    if err != nil {
        data["ERROR"] = fmt.Sprintf("exec chaincode err: %s", err)
        c.Data["json"] = data
        c.ServeJSON()
        return
    }
    data["SUCCESS"] = fmt.Sprintf("write key:[%s] with value:[%s]", key, value)
    c.Data["json"] = data
    c.ServeJSON()
}

func (c *MainController) QuerybykeyTest() {
	rData := &RVData{}
	status := &Status{}
	key := c.GetString("key")
	sdk, err := fabsdk.New(config.FromFile("./sdk_test.yaml"))
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to create new SDK: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(channelIDTest,
		fabsdk.WithUser(user),
		fabsdk.WithOrg(org))

	channelClient, err := channel.New(channelProvider)
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("create channel client fail: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}

	args := [][]byte{[]byte(key)}
	response, err := channelClient.Query(channel.Request{ChaincodeID: "sealtx", Fcn: "querybykey", Args: args})
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to query funds: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}
	status.Code = 200
	status.Msg = "success"
	rData.Status = status
	rData.Value = string(response.Payload)
	c.Data["json"] = rData
	c.ServeJSON()
}

func (c *MainController) HistoryTest() {
	rData := &RHData{}
	status := &Status{}
	key := c.GetString("key")
	sdk, err := fabsdk.New(config.FromFile("./sdk_test.yaml"))
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to create new SDK: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(channelIDTest,
		fabsdk.WithUser(user),
		fabsdk.WithOrg(org))

	channelClient, err := channel.New(channelProvider)
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("create channel client fail: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}

	args := [][]byte{[]byte(key)}
	response, err := channelClient.Query(channel.Request{ChaincodeID: "sealtx", Fcn: "history", Args: args})
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to query funds: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
	}
	_ = json.Unmarshal(response.Payload, &rData.History)
	status.Code = 200
	status.Msg = "success"
	rData.Status = status
	c.Data["json"] = rData
	c.ServeJSON()
}

func (c *MainController) GetTxIdSpecTest() {
	rData := &RIData{}
	status := &Status{}
	key := c.GetString("key")
	value := c.GetString("value")
	sdk, err := fabsdk.New(config.FromFile("./sdk_test.yaml"))
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to create new SDK: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(channelIDTest,
		fabsdk.WithUser(user),
		fabsdk.WithOrg(org))

	channelClient, err := channel.New(channelProvider)
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("create channel client fail: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}

	args := [][]byte{[]byte(key), []byte(value)}
	response, err := channelClient.Query(channel.Request{ChaincodeID: "sealtx", Fcn: "gettxidspec", Args: args})
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to query funds: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}
	status.Code = 200
	status.Msg = "success"
	rData.Status = status
	rData.TxId = string(response.Payload)
	c.Data["json"] = rData
	c.ServeJSON()
}

func (c *MainController) GetTxByIDTest() {
	rdata := &RTData{}
	status := &Status{}
	txInfo := &TxInfo{}
	blockInfo := &BlockInfo{}
	txid := c.GetString("txid")
	sdk, err := fabsdk.New(config.FromFile("./sdk_test.yaml"))
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to create new SDK: %s", err)
		rdata.Status = status
		c.Data["json"] = rdata
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(channelIDTest,
		fabsdk.WithUser(user),
		fabsdk.WithOrg(org))
	ledgerClient, err := ledger.New(channelProvider)
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf(fmt.Sprintf("create ledger client fail: %s", err))
		rdata.Status = status
		c.Data["json"] = rdata
		c.ServeJSON()
		return
	}

	block, err := ledgerClient.QueryBlockByTxID(fab.TransactionID(txid))
	if err != nil {
		status.Code = 404
		status.Msg = fmt.Sprintf(fmt.Sprintf("query block fail: %s", err))
		rdata.Status = status
		c.Data["json"] = rdata
		c.ServeJSON()
		return
	}
	blockInfo.Number = block.Header.Number
	blockInfo.CurrentHash = hex.EncodeToString(util.ComputeSHA256(tobytes(block.Header)))
	blockInfo.PreviousHash = hex.EncodeToString(block.Header.PreviousHash)
	txInfo.BlockInfo = blockInfo
	tx, err := ledgerClient.QueryTransaction(fab.TransactionID(txid))
	if err != nil {
		status.Code = 404
		status.Msg = fmt.Sprintf(fmt.Sprintf("query tx fail: %s", err))
		rdata.Status = status
		c.Data["json"] = rdata
		c.ServeJSON()
		return
	}
	payload := &cb.Payload{}
	if err = proto.Unmarshal(tx.TransactionEnvelope.Payload, payload); err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf(fmt.Sprintf("error reconstructing payload(%s)", err))
		rdata.Status = status
		c.Data["json"] = rdata
		c.ServeJSON()
		return
	}
	txInfo.parseEndorserTransaction(payload)
	// txActions := &fab_peer.TransactionAction{}
	// err = proto.Unmarshal(payload.Data, txActions)
	// if err != nil {
	// 	status.Code = 500
	// 	status.Msg = fmt.Sprintf(fmt.Sprintf("error reconstructing txActions(%s)", err))
	// 	rdata.Status = status
	// 	c.Data["json"] = rdata
	// 	c.ServeJSON()
	// 	return
	// }
	// _, ccAction, err := utils.GetPayloads(txActions)
	// if err != nil {
	// 	status.Code = 500
	// 	status.Msg = fmt.Sprintf(fmt.Sprintf("error reconstructing ccAction(%s)", err))
	// 	rdata.Status = status
	// 	c.Data["json"] = rdata
	// 	c.ServeJSON()
	// 	return
	// }
	// txInfo.RWSet = string(ccAction.Results)

	chhd := &cb.ChannelHeader{}
	err = proto.Unmarshal(payload.Header.ChannelHeader, chhd)
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf(fmt.Sprintf("error reconstructing channelheader(%s)", err))
		rdata.Status = status
		c.Data["json"] = rdata
		c.ServeJSON()
		return
	}
	txInfo.TxID = txid
	txInfo.ChannelId = chhd.GetChannelId()
	txInfo.ValidationCodeName = peer3.TxValidationCode_name[tx.GetValidationCode()]
	timestamp := &google_protobuf.Timestamp{}
	timestamp.Seconds = chhd.GetTimestamp().Seconds
	timestamp.Nanos = chhd.GetTimestamp().Nanos
	txInfo.Timestamp = timestamp
	rdata.TxInfo = txInfo
	status.Code = 200
	status.Msg = "success"
	rdata.Status = status
	c.Data["json"] = rdata
	c.ServeJSON()
}
