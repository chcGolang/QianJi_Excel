package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

const mysqlPath  = "root:123456@tcp(127.0.0.1:3307)/qianji_excel?charset=utf8"
type dataSource struct {
	MysqlDB *gorm.DB
}
var dataSourceDB *dataSource
var dataSourceLock sync.Mutex
func NewDataSource() *dataSource {
	dataSourceLock.Lock()
	defer dataSourceLock.Unlock()
	if dataSourceDB != nil{
		return dataSourceDB
	}
	db,err := initMysql()
	if err != nil{
		panic(err)
	}
	return &dataSource{
		MysqlDB:db,
	}

}

func initMysql()(db *gorm.DB,err error)  {
	db, err = gorm.Open("mysql", mysqlPath)
	if err != nil{
		return
	}
	if err = db.DB().Ping();err != nil{
		return
	}
	return

}
