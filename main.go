package main

import (
	"QianJi_Excel/alipay_excel"
	"QianJi_Excel/qianji_excel"
	"QianJi_Excel/utils"
	"QianJi_Excel/wechatpay_excel"
	"flag"
	"fmt"
	"log"
	"path"
	"strings"
	"time"
)

var(
	csvType int
	csvPath string
	downloadPath string
	downloadFilePath string
	oldPath = "./"

)
// 解析命令行参数
func initArgs()  {
	flag.IntVar(&csvType,"type",0,"excel类型(1:支付宝,2:微信)")
	flag.StringVar(&csvPath,"csvPath","","csv文件地址")
	flag.StringVar(&downloadPath,"downloadPath","","下载地址")
}

func main()  {
	var(
		err error
	)
	for{
		fmt.Println()
		fmt.Print("csv类型(1:支付宝,2:微信):")
		if _, err = fmt.Scanln(&csvType);err != nil{
			log.Printf("csv类型异常:%s \n",err.Error())
			continue
		}
		if csvType !=1 && csvType!=2{
			log.Printf("csv文件类型错误:%d \n",csvType)
			continue
		}

		fmt.Print("文件所在位置:")
		if _, err = fmt.Scanln(&csvPath);err != nil{
			log.Printf("文件所在位置异常:%s \n",err.Error())
			continue
		}

		fmt.Print("转换成功的文件存放位置:")

		if _, err = fmt.Scanln(&downloadPath);err != nil{
		}
		downloadPath = checkPath(downloadPath)
		oldPath = downloadPath
		switch csvType {
		case 1:
			if err := alipayCsv(csvPath);err != nil{
				log.Printf("解析异常:%s",err.Error())
				continue
			}
		case 2:
			if err = wechatPayCsv(csvPath);err != nil{
				log.Printf("解析异常:%s",err.Error())
				continue
			}
		}

		fmt.Println("文件存放位置:",downloadFilePath)
	}
}

func alipayCsv(filePath string)(err error){

	alipayDataList, err := alipay_excel.ReadAlipayExcel(filePath)
	if err != nil{
		return
	}
	excelMap := qianji_excel.AlipayDataToExcelMap(alipayDataList)
	downloadFilePath = downloadPath+"alipay_"+parseWithLocation()+".xls"
	err = utils.GetExcel().NewCategories(qianji_excel.ExcelQianJiKey).NewValueListMap(&excelMap).SaveAs(downloadFilePath)
	if err != nil{
		return
	}
	return
}

func wechatPayCsv(filePath string) (err error) {
	wechatPayDataList, err := wechatpay_excel.ReadWechatExcel(filePath)
	if err != nil{
		return
	}
	toExcelMap := qianji_excel.WechatPayDataToExcelMap(wechatPayDataList)
	downloadFilePath = downloadPath+"wechat_"+parseWithLocation()+".xls"
	err = utils.GetExcel().NewCategories(qianji_excel.ExcelQianJiKey).NewValueListMap(&toExcelMap).SaveAs(downloadFilePath)
	if err != nil{
		return
	}
	return
}


func checkPath(filePath string)(downloadPath string)  {

	if len(filePath) == 0{
		return oldPath
	}
	downloadPath = path.Clean(filePath)
	downloadPath = strings.Replace(downloadPath,"\\","/",-1)
	downloadPath+="/"
	return downloadPath
}



const TIME_LAYOUT = "2006-01-02_15-04-05"
func parseWithLocation() (timeStr string) {
	return time.Now().Format(TIME_LAYOUT)
}
