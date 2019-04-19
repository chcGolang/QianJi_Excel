package utils

import (
	"fmt"
	"reflect"
)

type reflectStruct struct {
	TypeOf reflect.Type
	ValueOf reflect.Value
}

func GetMapTagToStruct(structValue interface{})(dataMap map[string]string,err error)  {
	bol, reKin, reflectInfo := getReflectTypeAndValue(structValue, reflect.Struct)
	if !bol{
		err =  fmt.Errorf("structValue的类型应该:%s,不是:%s",reflect.Struct.String(),reKin.String())
		return
	}
	dataMap = make(map[string]string)
	numField := reflectInfo.TypeOf.NumField()
	for i:=0;i<numField;i++{
		field := reflectInfo.TypeOf.Field(i)
		jsonTag := field.Tag.Get("json")
		structName := field.Name
		valueStr := reflectInfo.ValueOf.FieldByName(structName).String()
		dataMap[jsonTag] = valueStr
	}
	return

}
// 获取指定类型的反射的Type和value
func getReflectTypeAndValue(value interface{},kindInt reflect.Kind)(bol bool,reKin reflect.Kind,reflectInfo reflectStruct)  {
	typeOf := reflect.TypeOf(value)
	valueOf := reflect.ValueOf(value)

	if typeOf.Kind() == reflect.Ptr{ //可能是指针
		typeOf = typeOf.Elem()
		valueOf = valueOf.Elem()

		bol = typeOf.Kind() == kindInt
		reKin = valueOf.Kind()
	}else {
		bol = typeOf.Kind() == kindInt
		reKin = valueOf.Kind()
	}
	reflectInfo = reflectStruct{
		TypeOf:typeOf,
		ValueOf:valueOf,
	}
	return
}
