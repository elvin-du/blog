package models

import "time"

const (
	ARTICLES_MODEL = iota
	INFO_MODEL
	ADMIN_MODEL
	COMMENT_MODEL
	VISITOR_MODEL
)

type Article struct {
	Id      int `orm:"index;unique;pk"`
	Title   string
	Content string
	CTime   time.Time `orm:"column(ctime)"`
	TagId   int       `orm:"column(tag_id)"`
}

func (this *Article) TableName() string {
	return "articles"
}

//build table index
func (this *Article) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}

type Comment struct {
	Id        int `orm:"index;unique"`
	Comment   string
	Ip        string
	CTime     time.Time `orm:"column(ctime)"`
	ArticleId int       `orm:"column(article_id)"`
}

func (this *Comment) TableName() string {
	return "comments"
}

type Tags struct {
	Id  int
	Tag string
}

type Admins struct {
	Id       int
	Name     string
	Password string
}
