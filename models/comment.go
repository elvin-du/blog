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

func (this *comment) CommentFromArticleId(id int) ([]orm.Params, error) {
	var maps []orm.Params
	_, err := orm.NewOrm().Raw("select * from comments where article_id= ?", id).Values(&maps)
	if nil != err {
		beego.Error(err)
		return nil, err
	}
	return maps, nil
}

func (this *comment) Latest() ([]orm.Params, error) {
	var maps []orm.Params
	_, err := orm.NewOrm().Raw("select * from comments order by ctime desc limit 10").Values(&maps)
	if nil != err {
		beego.Error(err)
		return maps, err
	}
	return maps, nil
}
