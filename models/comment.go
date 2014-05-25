package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
) // import your used driver

type comment byte

var _gComment comment = COMMENT_MODEL

func CommentModel() *comment {
	return &_gComment
}

func (this *comment) Add(id, comment, ip, nick, site, email string) error {
	o := orm.NewOrm()
	_, err := o.Raw("insert comments(ip,comment, ctime, article_id,nick,site,email) values(?,?,?,?,?,?,?)", ip, comment, time.Now(), id, nick, site, email).Exec()
	if nil != err {
		beego.Error(err)
		return err
	}
	return nil
}
