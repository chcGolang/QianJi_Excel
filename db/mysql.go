package db

import (
	"QianJi_Excel/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

const mysqlPath  = "root:123456@tcp(127.0.0.1:3307)/qianji_excel?charset=utf8&loc=Asia%2FShanghai"
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
	db.LogMode(true)
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
// 数据库panic异常回滚
func RecoverRollback(err *error,tx *gorm.DB)  {
	if r := recover(); r != nil {
		tx.Rollback()
		*err = utils.Errorf("数据连接异常")
	}
}
