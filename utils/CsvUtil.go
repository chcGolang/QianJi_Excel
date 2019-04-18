package utils

import (
	"bytes"
	"encoding/csv"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"strings"
)

// 解析csv文件数据
func ReadCsvToMap(bytesReader *bytes.Reader,keyLen int)[]map[string]interface{} {
	var(
		csvReader *csv.Reader
		titleKey []string
		dataMapList  = make([]map[string]interface{},0)
		dataMap map[string]interface{}
		err error
		record []string
	)
	csvReader = csv.NewReader(bytesReader)
	for{
		record, err = csvReader.Read()
		if err == io.EOF {
			// 读完退出
			break
		}
		if len(record) >= keyLen{
			if len(titleKey) == 0 { // 第一行为标题行
				titleKey = record
				continue
			}
			// 组装数据
			dataMap = assemblyMapData(titleKey,record)
			dataMapList = append(dataMapList,dataMap)
		}
	}
	return dataMapList
}

// 组装成map结构
func assemblyMapData(titleKey []string,dataInfo []string)map[string]interface{}  {
	var (
		dataMap = make(map[string]interface{})

	)
	var(
		key int
		value string
		replacer *strings.Replacer
	)
	for key,value = range titleKey{
		// 去除特殊字符
		value = trim(value)
		// 替换字符
		replacer = strings.NewReplacer("（", "(", "）", ")")
		value = replacer.Replace(value)

		if len(value) == 0{
			continue
		}
		dataMap[value] = trim(dataInfo[key])
	}
	return dataMap
}

// 去除特殊字符
func trim(dataStr string)string  {
	// 要去除的字符
	var cutsetSpecial = []string{" ","\t"}
	for _,value := range cutsetSpecial  {
		dataStr = strings.Trim(dataStr,value)
	}
	return dataStr
}

func Decode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
