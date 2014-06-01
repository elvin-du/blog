package models

import (
	"blog/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type articles byte

var _gArticles articles = ARTICLES_MODEL

func ArticlesModel() *articles {
	return &_gArticles
}

func (this *articles) ArticleFromId(id int) ([]orm.Params, error) {
	var maps []orm.Params
	_, err := orm.NewOrm().Raw("SELECT * FROM articles WHERE id = ?", id).Values(&maps)
	if nil != err {
		beego.Error(err)
		return nil, err
	}

	return maps, nil
}

func (this *articles) Count() (int64, error) {
	sqlStr := "select count(*) from articles"
	var maps []orm.Params
	_, err := orm.NewOrm().Raw(sqlStr).Values(&maps)
	if nil != err {
		beego.Error(err)
		return 0, err
	}
	beego.Debug(maps)

	totalStr := maps[0]["count(*)"].(string)
	total, err := strconv.ParseInt(totalStr, 10, 64)
	if nil != err {
		beego.Error(err)
		return 0, err
	}
	return total, nil
}

func (this *articles) Articles(num, page int64) ([]orm.Params, int64, error) {
	total, err := this.Count()
	if nil != err {
		return nil, 0, err
	}

	totalPage := total / num
	if 0 != total%num {
		totalPage++
	}

	var maps []orm.Params
	if page > totalPage {
		return maps, totalPage, nil
	}
	sqlStr := "SELECT a.id, title,excerpt, content, ctime,read_count, tag FROM articles a LEFT JOIN tags t ON a.tag_id = t.id ORDER BY ctime DESC LIMIT ?,?"
	_, err = orm.NewOrm().Raw(sqlStr, (page-1)*num, num).Values(&maps)
	if nil != err {
		beego.Error(err)
		return nil, totalPage, err
	}

	if len(maps) == 0 {
		return nil, totalPage, E_NOT_FOUND
	}
	return maps, totalPage, nil
}

func (this *articles) Add(title, content string) error {
	excerpt, content := utils.ExcerptContent(content)
	beego.Debug(excerpt)
	beego.Debug(content)
	o := orm.NewOrm()
	_, err := o.Raw("insert articles(title,excerpt,content, ctime) values(?,?,?,?)", title, excerpt, content, time.Now()).Exec()
	if nil != err {
		beego.Error(err)
		return err
	}
	return nil
}

func (this *articles) Del(id int) error {
	_, err := orm.NewOrm().Raw("delete from articles where id = ?", id).Exec()
	return err
}

func (this *articles) Update(id int, title, content string) error {
	excerpt, content := utils.ExcerptContent(content)
	beego.Debug(excerpt)
	beego.Debug(content)
	_, err := orm.NewOrm().Raw("update articles set title=?,excerpt=?, content=? where id = ?", title, excerpt, content, id).Exec()
	return err
}

func (this *articles) LogReadHistory(id int, ip string) error {
	_, err := orm.NewOrm().Raw("Insert read_history(article_id, ip,atime) values(?,?,?)", id, ip, time.Now()).Exec()
	if nil != err {
		beego.Error(err)
		return err
	}

	return nil
}
