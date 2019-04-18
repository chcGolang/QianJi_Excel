package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"path"
)

type excelStruct struct {
	sheet string
	// 标题
	categories []string
	// 数据
	values       *[]map[string]interface{}
	excelizeFile *excelize.File
	// 开始行数
	lineNumber int
	// 是否需要标题
	titleBool bool
}

var excelKey = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG"}

// 创建excel的方式导入excel
func GetExcel() *excelStruct {
	return &excelStruct{
		sheet:        "Sheet1",
		excelizeFile: excelize.NewFile(),
		lineNumber:   2,
		titleBool:    true,
	}
}

// 打开excel的方式导入excel(追加的方式)
func GetOpenExcel(dirPath, fileName string) (*excelStruct, error) {
	var (
		err          error
		excelizeFile *excelize.File
		filePath     string
	)
	filePath = dirPath + fileName
	if err = createDir(dirPath); err != nil {
		return nil, err
	}
	filePath = dirPath + fileName
	if !exists(filePath) {
		excelizeFile = excelize.NewFile()
		excelizeFile.Path = filePath
	} else {
		excelizeFile, err = excelize.OpenFile(filePath)
	}
	return &excelStruct{
		sheet:        "Sheet1",
		excelizeFile: excelizeFile,
		lineNumber:   2,
		titleBool:    true,
	}, err
}

func createDir(dirPath string) error {
	var err error
	if !exists(path.Dir(dirPath)) {
		if err = os.MkdirAll(path.Dir(dirPath), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 设置数据开始行数,默认为2
func (e *excelStruct) SetStartLineNum(startLineNum int) *excelStruct {
	if startLineNum < 1 {
		startLineNum = 1
	}
	e.lineNumber = startLineNum
	return e
}

// 设置是否需要标题
func (e *excelStruct) SetTitleBool(titleBool bool) *excelStruct {
	e.titleBool = titleBool
	return e
}

func (e *excelStruct) NewSheet(sheet string) *excelStruct {
	e.sheet = sheet
	return e
}

func (e *excelStruct) NewCategories(categories []string) *excelStruct {
	e.categories = categories
	return e
}

func (e *excelStruct) NewValueListMap(values *[]map[string]interface{}) *excelStruct {
	e.values = values
	return e
}

func (e *excelStruct) SaveAs(filePath string) error {
	e.assembleData()
	return e.excelizeFile.SaveAs(filePath)
}

func (e *excelStruct) Save() error {
	e.assembleData()
	return e.excelizeFile.Save()
}

func (e *excelStruct) Write(w io.Writer) error {
	e.assembleData()
	return e.excelizeFile.Write(w)
}

// 组装数据
func (e *excelStruct) assembleData() {
	values := e.values
	excelizeFile := e.excelizeFile
	sheet := e.sheet
	categories := e.categories

	if e.titleBool {
		for key, value := range categories {
			excelizeFile.SetCellValue(sheet, getExcelKey(excelKey[key], 1), value)
		}
	}

	cellky := e.lineNumber
	cellKey := 0
	if values != nil {
		for key, value := range *values {
			cellKey = key + cellky
			for i := 0; i < len(categories); i++ {
				mapkey := categories[i]
				excelizeFile.SetCellValue(sheet, getExcelKey(excelKey[i], cellKey), value[mapkey])
			}

		}
	}

}

func getExcelKey(key string, intkey int) string {
	return fmt.Sprintf("%s%d", key, intkey)
}

type readExcelStruct struct {
	// 开始行数
	startLint int
	// 标题名称
	titleKey []string
	// 导出的数据
	dataValue []map[string]interface{}
}
