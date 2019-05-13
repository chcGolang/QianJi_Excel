package dao

import (
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"github.com/jinzhu/gorm"
)

type ICsvFileInfoDao interface {
	// 添加
	Save(csvFileInfo *model.CsvFileInfo)(err error)

}

type csvFileInfoDao struct {
	db *gorm.DB
}


func (c *csvFileInfoDao) Save(csvFileInfo *model.CsvFileInfo) (err error) {
	return c.db.Save(csvFileInfo).Error
}

func NewCsvFileInfoDao()ICsvFileInfoDao  {
	return csvFileInfoDaoImpl
}
var csvFileInfoDaoImpl = &csvFileInfoDao{
	db:db.NewDataSource().MysqlDB,
}

