package routers

import (
	"bx/controllers"
	"bx/controllers/user"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/auth/toLogin", &user.AuthController{})
}
