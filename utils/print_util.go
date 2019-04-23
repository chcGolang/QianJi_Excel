package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(dataInfo interface{})  {
	if dataInfo==nil{
		fmt.Println("nil")
	}
	bytes, _ := json.Marshal(dataInfo)
	fmt.Println(string(bytes))
}
