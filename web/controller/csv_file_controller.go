package controller

import (
	"QianJi_Excel/service"
	"QianJi_Excel/vo"
	"fmt"
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
	file_key := c.Ctx.URLParam("file_key")
	fmt.Println(file_key)
	id := 0
	if id,err = c.CsvFileInfoService.UploadCsvFile(file, header,file_key);err != nil{
		return failError(err,fail_service)
	}

	return succeesData(id)
}

// 添加文件信息
func (c *CsvFileController)PostFile_info()resultResp  {
	voInfo :=  vo.CsvFileSaveInfo{}
	if err := c.Ctx.ReadJSON(&voInfo);err != nil{
		return failError(err,fali_param)
	}
	if err := c.CsvFileInfoService.SaveCsvFileInfo(voInfo);err != nil{
		return failError(err,fail_service)
	}

	return succeesData("")
}
