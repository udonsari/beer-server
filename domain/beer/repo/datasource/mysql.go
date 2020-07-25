package datasource

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func NewMySQLDataSource(dataSourceName string) orm.Ormer {
	o := orm.NewOrm()

	// register model
	orm.RegisterModel(new(DBBeer))
	// set default database
	// 포맷 재확인
	orm.RegisterDataBase("default", "mysql", dataSourceName, 30)
	// create table
	orm.RunSyncdb("default", false, true)

	return o
}
