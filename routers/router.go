package routers

import (
	"explorer-service/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/querybykey", &controllers.MainController{}, "get:Querybykey")
	beego.Router("/querybykey/test", &controllers.MainController{}, "get:QuerybykeyTest")
	beego.Router("/history", &controllers.MainController{}, "get:History")
	beego.Router("/history/test", &controllers.MainController{}, "get:HistoryTest")
	beego.Router("/gettxidspec", &controllers.MainController{}, "get:GetTxIdSpec")
	beego.Router("/gettxidspec/test", &controllers.MainController{}, "get:GetTxIdSpecTest")
	beego.Router("/gettxbyid", &controllers.MainController{}, "get:GetTxByID")
	beego.Router("/gettxbyid/test", &controllers.MainController{}, "get:GetTxByIDTest")
	beego.Router("/seal/test", &controllers.MainController{}, "get:SealTest")
}
