package user

/**
  登录控制器
*/

import (
	"bx/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) Get() {
	this.Data["Website"] = "99999999999999"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "auth/login.tpl"
}

func (this *AuthController) Post() {
	o := orm.NewOrm()
	u := new(models.User)
	u.Name = "straw"
	u.PassWord = "123"
	o.Insert(u)
}
