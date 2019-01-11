package controllers

import (
	"github.com/astaxie/beego"
)

var (
	channelID = beego.AppConfig.String("channelid")
	channelIDTest = beego.AppConfig.String("channelidtest")
	user      = beego.AppConfig.String("user")
	org       = beego.AppConfig.String("org")
)
