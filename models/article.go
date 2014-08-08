package models

import (
	"blog/utils"
	"strconv"
	"strings"
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
	sqlStr := "SELECT * FROM articles ORDER BY ctime DESC LIMIT ?,?"
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

func (this *articles) Add(title, content string, tags []string) error {
	excerpt, content := utils.ExcerptContent(content)
	o := orm.NewOrm()
	result, err := o.Raw("insert articles(title,excerpt,content, ctime) values(?,?,?,?)", title, excerpt, content, time.Now()).Exec()
	if nil != err {
		beego.Error(err)
		return err
	}

	articleId, err := result.LastInsertId()
	if nil != err {
		beego.Error(err)
		return err
	}

	var ts = []string{}
	if len(tags) == 0 {
		ts = append(ts, "其他")
	} else {
		if tags[0] == "" {
			ts = append(ts, "其他")
		} else {
			ts = tags
		}
	}

	for _, v := range ts {
		var maps []orm.Params
		_, err := orm.NewOrm().Raw("select id from tags where tag=?", v).Values(&maps)
		if nil != err {
			beego.Error(err)
			return err
		}

		var tagId int64 = 0
		if 0 == len(maps) {
			result, err = orm.NewOrm().Raw("insert tags(tag) values(?)", v).Exec()
			if nil != err {
				beego.Error(err)
				return err
			}
			tagId, err = result.LastInsertId()
			if nil != err {
				beego.Error(err)
				return err
			}
		} else {
			tagIdStr := maps[0]["id"].(string)
			tagId, err = strconv.ParseInt(tagIdStr, 10, 64)
			if nil != err {
				beego.Error(err)
				return err
			}
		}

		_, err = orm.NewOrm().Raw("insert article_tag_relation(article_id,tag_id) values(?,?)", articleId, tagId).Exec()
		if nil != err {
			beego.Error(err)
			return err
		}
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

func (this *articles) Hot() ([]orm.Params, error) {
	var maps []orm.Params
	_, err := orm.NewOrm().Raw("select * from articles order by read_count desc limit 10").Values(&maps)
	if nil != err {
		beego.Error(err)
		return maps, err

	}
	return maps, nil
}

func (this *articles) Latest() ([]orm.Params, error) {
	var maps []orm.Params
	_, err := orm.NewOrm().Raw("select * from articles order by ctime desc limit 10").Values(&maps)
	if nil != err {
		beego.Error(err)
		return maps, err

	}
	return maps, nil
}

func (this *articles) Tags(id int64) ([]string, error) {
	var tags = []string{}

	var maps []orm.Params
	_, err := orm.NewOrm().Raw("select tag_id from article_tag_relation where article_id=?", id).Values(&maps)
	if nil != err {
		beego.Error(err)
		return tags, err
	}

	for _, v := range maps {
		idStr := v["tag_id"].(string)
		tag_id, err := strconv.ParseInt(idStr, 10, 64)
		if nil != err {
			beego.Error(err)
			return tags, err
		}

		var maps2 = []orm.Params{}
		_, err = orm.NewOrm().Raw("select tag from tags where id=?", tag_id).Values(&maps2)
		if nil != err {
			beego.Error(err)
			return tags, err
		}
		if 0 != len(maps2) {
			tags = append(tags, maps2[0]["tag"].(string))
		}
	}

	return tags, nil
}

func (this *articles) ArticlesFromTagId(id int64) ([]orm.Params, error) {
	var maps []orm.Params
	_, err := orm.NewOrm().Raw("select * from article_tag_relation where tag_id=?", id).Values(&maps)
	if nil != err {
		beego.Error(err)
		return maps, err

	}

	var in []string
	for _, v := range maps {
		in = append(in, v["article_id"].(string))
	}
	sqlIn := strings.Join(in, ",")
	sqlStr := "select * from articles where id in(" + sqlIn + ")"
	beego.Debug("sql:", sqlStr)
	var maps2 []orm.Params
	_, err = orm.NewOrm().Raw(sqlStr).Values(&maps2)
	if nil != err {
		beego.Error(err)
		return maps, err

	}
	return maps2, nil
}

func (this *articles) Search(key string) ([]orm.Params, error) {
	var maps []orm.Params
	sqlStr := `SELECT * FROM articles WHERE CONCAT(UPPER(excerpt),UPPER(content)) LIKE BINARY CONCAT('%',UPPER('` + key + `'),'%')`
	beego.Debug("sql:", sqlStr)
	_, err := orm.NewOrm().Raw(sqlStr).Values(&maps)
	if nil != err {
		beego.Error(err)
		return nil, err
	}

	return maps, nil
}
