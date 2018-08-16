package routers

import (
	"github.com/astaxie/beego"
	"tiger-blockchain-api/controllers"
)

func init() {
	beego.Router("/querybykey", &controllers.MainController{}, "get:Querybykey")
	beego.Router("/history", &controllers.MainController{}, "get:History")
}
