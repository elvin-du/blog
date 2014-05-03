package models

import _ "github.com/go-sql-driver/mysql" // import your used driver

type admin byte

var _gAdmin admin = ADMIN_MODEL

func AdminModel() *admin {
	return &_gAdmin
}
