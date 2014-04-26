package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	// register model
	orm.RegisterModel(new(Article))

	// set default database
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:oliver@/blog?charset=utf8", maxIdle, maxConn)
}

type articles byte

var _gArticles articles = ARTICLES_MODEL

func ArticlesModel() *articles {
	return &_gArticles
}

func (this *articles) ArticleFromId(id int) (*Article, error) {
	o := orm.NewOrm()

	a := Article{Id: id}
	err := o.Read(&a)
	if nil != err {
		beego.Error(err)
		return nil, err
	}

	return &a, nil
}
