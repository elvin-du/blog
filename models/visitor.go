package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// import your used driver

type visitor byte

var _gVisitor visitor = VISITOR_MODEL

func VisitorModel() *visitor {
	return &_gVisitor
}

func (this *visitor) Log(ip string, atime time.Time) error {
	_, err := orm.NewOrm().Raw("insert visitors(ip,atime) values(?,?)", ip, atime).Exec()

	if nil != err {
		beego.Error(err)
		return err
	}
	return nil
}
