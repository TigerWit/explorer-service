package controllers

import (
	"encoding/asn1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	peer3 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/utils"
	"math"
)

type MainController struct {
	beego.Controller
}

type RVData struct {
	Status *Status `json:"status"`
	Value  string  `json:"value"`
}

func (c *MainController) Querybykey() {
	rData := &RVData{}
	status := &Status{}
	key := c.GetString("key")
	sdk, err := fabsdk.New(config.FromFile("./sdk.yaml"))
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to create new SDK: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(channelID,
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

type RHData struct {
	Status  *Status  `json:"status"`
	History []string `json:"history"`
}

func (c *MainController) History() {
	rData := &RHData{}
	status := &Status{}
	key := c.GetString("key")
	sdk, err := fabsdk.New(config.FromFile("./sdk.yaml"))
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to create new SDK: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(channelID,
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

type RIData struct {
	Status *Status `json:"status"`
	TxId   string  `json:"tx_id"`
}

func (c *MainController) GetTxIdSpec() {
	rData := &RIData{}
	status := &Status{}
	key := c.GetString("key")
	value := c.GetString("value")
	sdk, err := fabsdk.New(config.FromFile("./sdk.yaml"))
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to create new SDK: %s", err)
		rData.Status = status
		c.Data["json"] = rData
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(channelID,
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

type asn1Header struct {
	Number       int64
	PreviousHash []byte
	DataHash     []byte
}

func tobytes(cb *common.BlockHeader) []byte {
	asn1Header := asn1Header{
		PreviousHash: cb.PreviousHash,
		DataHash:     cb.DataHash,
	}
	if cb.Number > uint64(math.MaxInt64) {
		panic(fmt.Errorf("Golang does not currently support encoding uint64 to asn1"))
	} else {
		asn1Header.Number = int64(cb.Number)
	}
	result, err := asn1.Marshal(asn1Header)
	if err != nil {
		panic(err)
	}
	return result
}

type RTData struct {
	Status *Status `json:"status"`
	TxInfo *TxInfo `json:"tx_info"`
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type TxInfo struct {
	TxID               string                     `json:"tx_id"`
	Key                string                     `json:"key"`
	Value              string                     `json:"value"`
	ChannelId          string                     `json:"channel_id"`
	BlockInfo          *BlockInfo                 `json:"block_info"`
	ValidationCodeName string                     `json:"validation_code_name"`
	Timestamp          *google_protobuf.Timestamp `json:"timestamp"`
}

type BlockInfo struct {
	Number       uint64 `json:"number"`
	CurrentHash  string `json:"current_hash"`
	PreviousHash string `json:"previous_hash"`
}

func (c *MainController) GetTxByID() {
	rdata := &RTData{}
	status := &Status{}
	txInfo := &TxInfo{}
	blockInfo := &BlockInfo{}
	txid := c.GetString("txid")
	sdk, err := fabsdk.New(config.FromFile("./sdk.yaml"))
	if err != nil {
		status.Code = 500
		status.Msg = fmt.Sprintf("Failed to create new SDK: %s", err)
		rdata.Status = status
		c.Data["json"] = rdata
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(channelID,
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

func (txInfo *TxInfo) parseEndorserTransaction(payload *cb.Payload) error {
	var err error
	var tx *peer.Transaction
	if tx, err = utils.GetTransaction(payload.Data); err != nil {
		return err
	}

	fmt.Printf("        actions\n")
	for _, action := range tx.Actions {
		// var capayload *peer.ChaincodeActionPayload
		var ca *peer.ChaincodeAction
		if _, ca, err = utils.GetPayloads(action); err != nil {
			return err
		}

		// fmt.Printf("            endorsers\n")
		// for _, endorser := range capayload.Action.Endorsements {
		// 	var mspid, subject string
		// 	if mspid, subject, err = decodeSerializedIdentity(endorser.Endorser); err != nil {
		// 		return err
		// 	}
		// 	fmt.Printf("                endorser[%s:%s]\n", mspid, subject)
		// }

		fmt.Printf("            RWSet\n")
		txRWSet := &rwsetutil.TxRwSet{}
		if err = txRWSet.FromProtoBytes(ca.Results); err != nil {
			return err
		}
		for _, nsRWSet := range txRWSet.NsRwSets {
			ns := nsRWSet.NameSpace
			if ns != "lscc" { // skip system chaincode
				fmt.Printf("                ns=[%v]\n", ns)
				fmt.Printf("                RDSet\n")
				for _, kvRead := range nsRWSet.KvRwSet.Reads {
					fmt.Printf("                     key=[%v]\n", kvRead.Key)
				}
				fmt.Printf("                WRSet\n")
				for _, kvWrite := range nsRWSet.KvRwSet.Writes {
					if kvWrite.IsDelete {
						fmt.Printf("                     key=[%v] op=[delete]\n", kvWrite.Key)
					} else {
						fmt.Printf("                     key=[%v] op=[write]\n", kvWrite.Key)
						txInfo.Key = kvWrite.Key
						txInfo.Value = string(kvWrite.Value)
					}
				}
			}
		}
	}
	return nil
}
