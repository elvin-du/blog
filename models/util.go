package models

import "github.com/astaxie/beego/orm"

func init() {
	// set default database
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:oliver@/blog?charset=utf8", maxIdle, maxConn)
}
