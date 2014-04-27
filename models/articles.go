package models

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
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

func (this *articles) ArticleFromId(id int) ([]orm.Params, error) {
	o := orm.NewOrm()
	sqlStr := "SELECT title,content,a.ctime a_ctime,comment,ip,c.ctime FROM articles a ,comments c WHERE a.id = c.article_id AND a.id = " + strconv.Itoa(id)
	var maps []orm.Params
	_, err := o.Raw(sqlStr).Values(&maps)
	if nil != err {
		beego.Error(err)
		return nil, err
	}
	beego.Debug(maps)

	return maps, nil
}
