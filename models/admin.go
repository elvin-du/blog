package models

import (
	"blog/utils"
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type admin byte

var _gAdmin admin = ADMIN_MODEL

func AdminModel() *admin {
	return &_gAdmin
}

func (this *admin) Login(name, passwd string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(passwd))
	if nil != err {
		beego.Error(err)
		return "", err
	}
	passwd = hex.EncodeToString(h.Sum(nil))

	var maps []orm.Params
	sqlStr := "select id from admins where name = ? and password = ?"
	_, err = orm.NewOrm().Raw(sqlStr, name, passwd).Values(&maps)
	if nil != err {
		beego.Error(err)
		return "", err
	}

	if len(maps) < 1 {
		return "", E_NOT_FOUND
	}
	idi, ok := maps[0]["id"]
	if !ok {
		return "", E_NOT_FOUND
	}

	idStr, ok := idi.(string)
	if !ok {
		return "", E_NOT_FOUND
	}

	id, err := strconv.Atoi(idStr)
	if nil != err {
		beego.Error(err)
		return "", err
	}

	return this.Session(id)
}

func (this *admin) Session(uid int) (string, error) {
	session := utils.UUID()
	sqlStr := "insert admin_sessions(uid,session) values(?,?)"
	_, err := orm.NewOrm().Raw(sqlStr, uid, session).Exec()
	if nil != err {
		beego.Error(err)
		return "", err
	}

	return session, nil
}

func (this *admin) DelSession(session string) error {
	sqlStr := "delete from admin_sessions where session= ?"
	_, err := orm.NewOrm().Raw(sqlStr, session).Exec()
	if nil != err {
		beego.Error(err)
		return err
	}
	return nil
}

func (this *admin) ValidSession(sess string) bool {
	sqlStr := "select id from admin_sessions where session = ?"
	var maps []orm.Params
	_, err := orm.NewOrm().Raw(sqlStr, sess).Values(&maps)
	if nil != err {
		beego.Error(err)
		return false
	}

	if len(maps) == 1 {
		return true
	}
	return false
}
