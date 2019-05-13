package service

import (
	"QianJi_Excel/dao"
	"QianJi_Excel/model"
	"QianJi_Excel/vo"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type ICsvFileInfoService interface {
	// 上传文件
	UploadCsvFile(file multipart.File,fileHeader *multipart.FileHeader,file_key string)(id int,err error)
	// 添加文件信息
	SaveCsvFileInfo(voInfo vo.CsvFileSaveInfo)(err error)
}

func NewCsvFileInfoService()ICsvFileInfoService  {
	return csvFileInfoServiceImpl
}
var csvFileInfoServiceImpl = &csvFileInfoService{
	CsvFileInfoDao:dao.NewCsvFileInfoDao(),
}
type csvFileInfoService struct {
	CsvFileInfoDao dao.ICsvFileInfoDao
}

func (c *csvFileInfoService) SaveCsvFileInfo(voInfo vo.CsvFileSaveInfo) (err error) {
	billTime, err := time.Parse("2006-01", voInfo.BillTime)
	if err != nil{
		return
	}
	info := model.CsvFileInfo{
		BillTime:&billTime,
		FileKey:voInfo.FileKey,
		CsvType:voInfo.CsvType,
		Status:1,
	}
	return c.CsvFileInfoDao.UpdateByFileKey(info)

}

func (c *csvFileInfoService) UploadCsvFile(file multipart.File, fileHeader *multipart.FileHeader,file_key string) (id int,err error) {
	filename := fileHeader.Filename
	filePath := "./"+filename
	out, err := os.OpenFile("./"+filename,
		os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	defer out.Close()
	if err != nil{
		return
	}
	if _, err = io.Copy(out, file);err != nil{
		return
	}
	info := model.CsvFileInfo{
		UploadFilePath: filePath,
		Status:         0,
		BillTime:nil,
		FileKey:file_key,
	}
	if err = c.CsvFileInfoDao.Save(&info);err != nil{
		out.Close()
		os.Remove(filePath)

		return
	}
	id = info.Id
	return
}



