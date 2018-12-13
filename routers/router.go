package routers

import (
	"explorer-service/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/querybykey", &controllers.MainController{}, "get:Querybykey")
	beego.Router("/history", &controllers.MainController{}, "get:History")
	beego.Router("/gettxidspec", &controllers.MainController{}, "get:GetTxIdSpec")
	beego.Router("/gettxbyid", &controllers.MainController{}, "get:GetTxByID")
}
