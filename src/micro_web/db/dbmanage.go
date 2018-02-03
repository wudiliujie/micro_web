package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"micro_web/core"
)
type DBManage struct {
	db *sqlx.DB
}
var G_dbmanage DBManage

func (manage * DBManage)Init()  {
	var err error
	manage.db, err = sqlx.Connect("mysql", "root:111111@tcp(127.0.0.1:3306)/mytest?charset=utf8&parseTime=true")
	if err != nil {
		core.Check(err)
	}
}
func (manage *DBManage)GetDB()(db* sqlx.DB)  {
	return manage.db
}
