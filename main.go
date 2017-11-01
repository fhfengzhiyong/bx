package main

import (
	_ "bx/routers"
	_ "bx/system"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
