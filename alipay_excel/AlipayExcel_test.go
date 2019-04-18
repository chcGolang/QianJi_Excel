package alipay_excel

import (
	"fmt"
	"testing"
)

func TestReadAlipayExcel(t *testing.T) {
	alipayDataList, err := ReadAlipayExcel("./alipay_record_20190416_1004_1.csv")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%+v",alipayDataList)
}
