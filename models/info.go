package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type info byte

var _gInfo info = INFO_MODEL

func InfoModel() *info {
	return &_gInfo
}

func (this *info) InfoFromName(name string) ([]orm.Params, error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("select * from infos where name = '%s'", name)
	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if nil != err {
		beego.Error(err)
		return nil, err
	}

	beego.Debug(maps)
	if len(maps) == 0 {
		return nil, E_NOT_FOUND
	}
	return maps, nil
}

func (this *info) UdateAbout(update string) error {
	_, err := orm.NewOrm().Raw(`update infos set value=? where name='about'`, update).Exec()
	if nil != err {
		beego.Error(err)
		return err
	}
	return nil
}
