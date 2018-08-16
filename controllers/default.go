package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type MainController struct {
	beego.Controller
}

const (
	CHANNEL_ID = "" //channel名称
	ORG_NAME   = "" //组织名称
	USER       = "" //组织用户身份
)

func (c *MainController) Querybykey() {
	key := c.GetString("key")
	data := make(map[string]string)
	sdk, err := fabsdk.New(config.FromFile("./sdk.yaml"))
	if err != nil {
		data["ERROR"] = fmt.Sprintf("Failed to create new SDK: %s", err)
		c.Data["json"] = data
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(CHANNEL_ID,
		fabsdk.WithUser(USER),
		fabsdk.WithOrg(ORG_NAME))

	channelClient, err := channel.New(channelProvider)
	if err != nil {
		data["ERROR"] = fmt.Sprintf("create channel client fail: %s", err)
		c.Data["json"] = data
		c.ServeJSON()
		return
	}

	args := [][]byte{[]byte(key)}
	response, err := channelClient.Query(channel.Request{ChaincodeID: "sealtx", Fcn: "querybykey", Args: args})
	if err != nil {
		data["ERROR"] = fmt.Sprintf("Failed to query funds: %s", err)
		c.Data["json"] = data
		c.ServeJSON()
		return
	}
	c.Data["json"] = string(response.Payload)
	c.ServeJSON()
}

func (c *MainController) History() {
	key := c.GetString("key")
	data := make(map[string]string)
	sdk, err := fabsdk.New(config.FromFile("./sdk.yaml"))
	if err != nil {
		data["ERROR"] = fmt.Sprintf("Failed to create new SDK: %s", err)
		c.Data["json"] = data
		c.ServeJSON()
		return
	}
	defer sdk.Close()

	channelProvider := sdk.ChannelContext(CHANNEL_ID,
		fabsdk.WithUser(USER),
		fabsdk.WithOrg(ORG_NAME))

	channelClient, err := channel.New(channelProvider)
	if err != nil {
		data["ERROR"] = fmt.Sprintf("create channel client fail: %s", err)
		c.Data["json"] = data
		c.ServeJSON()
		return
	}

	args := [][]byte{[]byte(key)}
	response, err := channelClient.Query(channel.Request{ChaincodeID: "sealtx", Fcn: "history", Args: args})
	if err != nil {
		data["ERROR"] = fmt.Sprintf("Failed to query funds: %s", err)
		c.Data["json"] = data
		c.ServeJSON()
	}
	result := []string{}
	_ = json.Unmarshal(response.Payload, &result)
	c.Data["json"] = result
	c.ServeJSON()
}
