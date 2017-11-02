package system

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbuser string = "root"
	dbpwd  string = "1234"
	dbname string = "104.128.82.190:3306/bx"
)

func init() {
	conn := dbuser + ":" + dbpwd + "@" + dbname + "?charset=utf8" //组合成连接串
	orm.RegisterDriver("mysql", orm.DRMySQL)                      //注册mysql驱动
	orm.RegisterDataBase("default", "mysql", conn)                //设置conn中的数据库为默认使用数据库
	orm.RunSyncdb("default", false, false)                        //后一个使用true会带上很多打印信息，数据库操作和建表操作的；第二个为true代表强制创建表
}
