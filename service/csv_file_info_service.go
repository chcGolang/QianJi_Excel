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
	UploadCsvFile(file multipart.File,fileHeader *multipart.FileHeader,fileUpload vo.FileUpload)(id int,err error)
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


func (c *csvFileInfoService) UploadCsvFile(file multipart.File, fileHeader *multipart.FileHeader,fileUpload vo.FileUpload) (id int,err error) {
	billTime, err := time.Parse("2006-01", fileUpload.BillTime)
	if err != nil{
		return
	}
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
		Status:         1,
		BillTime:&billTime,
		CsvType:fileUpload.CsvType,
	}
	if err = c.CsvFileInfoDao.Save(&info);err != nil{
		out.Close()
		os.Remove(filePath)

		return
	}
	id = info.Id
	return
}



