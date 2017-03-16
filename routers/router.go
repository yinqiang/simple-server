package routers

import (
	"github.com/yinqiang/simple-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
