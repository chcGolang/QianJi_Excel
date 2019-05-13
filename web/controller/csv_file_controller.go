package controller

import (
	"QianJi_Excel/service"
	"QianJi_Excel/vo"
	"github.com/kataras/iris"
)

type CsvFileController struct {
	Ctx iris.Context
	CsvFileInfoService service.ICsvFileInfoService
}

// 上传文件
func (c *CsvFileController)PostUpload()resultResp  {
	file, header, err := c.Ctx.FormFile("csv_file")
	if err != nil{
		return failError(err,fali_param)
	}
	upload := vo.FileUpload{}
	if err = c.Ctx.ReadForm(&upload);err != nil{
		return failError(err,fali_param)
	}

	id := 0
	if id,err = c.CsvFileInfoService.UploadCsvFile(file, header,upload);err != nil{
		return failError(err,fail_service)
	}

	return succeesData(id)
}


