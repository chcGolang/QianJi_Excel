package dao

import (
	"QianJi_Excel/db"
	"QianJi_Excel/model"
	"github.com/jinzhu/gorm"
)

type ICsvFileInfoDao interface {
	// 添加
	Save(csvFileInfo *model.CsvFileInfo)(err error)
	// 根据filekey更新
	UpdateByFileKey(modelInfo model.CsvFileInfo)(err error)
}

type csvFileInfoDao struct {
	db *gorm.DB
}

func (c *csvFileInfoDao) UpdateByFileKey(modelInfo model.CsvFileInfo) (err error) {
	mapUpdate := map[string]interface{}{
		"bill_time":*modelInfo.BillTime,
		"status":modelInfo.Status,
		"csv_type":modelInfo.CsvType,
	}
	return c.db.Table(model.CsvFileInfo{}.TableName()).Where("file_key = ?",modelInfo.FileKey).Updates(mapUpdate).Error
}

func (c *csvFileInfoDao) Save(csvFileInfo *model.CsvFileInfo) (err error) {
	if err = c.db.Delete(model.CsvFileInfo{}, "file_key = ?", csvFileInfo.FileKey).Error;err != nil{
		return
	}
	return c.db.Save(csvFileInfo).Error
}

func NewCsvFileInfoDao()ICsvFileInfoDao  {
	return csvFileInfoDaoImpl
}
var csvFileInfoDaoImpl = &csvFileInfoDao{
	db:db.NewDataSource().MysqlDB,
}

