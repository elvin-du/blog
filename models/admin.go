package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type admin byte

var _gAdmin admin = ADMIN_MODEL

func AdminModel() *admin {
	return &_gAdmin
}

//@RET:cookie
func (this *admin) Login(name, passwd string) error {
	h := md5.New()
	_, err := h.Write([]byte(passwd))
	if nil != err {
		beego.Error(err)
		return err
	}
	passwd = hex.EncodeToString(h.Sum(nil))

	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("select count(1) from blog.admins where name = ? and password = ?", name, passwd).Values(&maps)
	if nil != err {
		beego.Error(err)
		return err
	}
	if num <= 0 {
		return E_NOT_FOUND
	}

	return nil
}
