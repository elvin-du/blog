package models

import "time"

const (
	ARTICLES_MODEL = iota
)

type Article struct {
	Id      int `orm:"index;unique;pk"`
	Title   string
	Content string
	CTime   time.Time `orm:"column(ctime)", json:"ctime"`
	TagId   int       `orm:"column(tag_id)", json:"tag_id"`
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
	Id        int
	Comment   string
	Ip        string
	CTime     time.Time `json:"ctime"`
	ArticleId int       `json:"article_id"`
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
