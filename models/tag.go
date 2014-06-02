package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type tag byte

var _gTag tag = TAG_MODEL

func TagModel() *tag {
	return &_gTag
}

func (this *tag) Tags() ([]orm.Params, error) {
	var maps []orm.Params
	_, err := orm.NewOrm().Raw("select * from tags").Values(&maps)
	if nil != err {
		beego.Error(err)
		return nil, err
	}
	return maps, nil
}
