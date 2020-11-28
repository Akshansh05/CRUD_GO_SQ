package config

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//root:root@/simplerest?charset=utf8&parseTime=True&loc=Local
// “user:password@tcp(Hostname:Port)/dbname?charset=utf8&parseTime=True&loc=Local”
const (
	ConstConnectionString = "root:root@tcp(127.0.0.1:3306)/local?charset=utf8&parseTime=True"
	ConstAlias            = "default"
	ConstDriver           = "mysql"
)

//register orm to the connected database of which the connection string in given
func Connect() {

	//check if db is already registered, if not register it
	_, err := orm.GetDB(ConstAlias)
	if err != nil {
		errInRegister := orm.RegisterDataBase(ConstAlias, ConstDriver, ConstConnectionString)
		if errInRegister != nil {
			panic(errInRegister)
		}
	}
}
