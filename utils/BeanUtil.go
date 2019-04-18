package utils

import (
	"encoding/gob"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"reflect"
	"fmt"
	"strconv"
)
// 深度对象拷贝
// param &dst新对象
// param src原对象
func DeepCopy(dst interface{}, src interface{} ) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// json格式化
func Marshal(v interface{}) ([]byte, error){
	return json.Marshal(v)
}

// json转换实体
func Unmarshal(data []byte, v interface{}) error   {
	return json.Unmarshal(data,v)
}


func setField(obj interface{}, name string, value interface{}) error {
	structData := reflect.ValueOf(obj).Elem()
	fieldValue := structData.FieldByName(name)

	if !fieldValue.IsValid() {
		return fmt.Errorf("utils.setField() No such field: %s in obj ", name)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value ", name)
	}

	fieldType := fieldValue.Type()
	val       := reflect.ValueOf(value)

	valTypeStr   := val.Type().String()
	fieldTypeStr := fieldType.String()
	if valTypeStr == "float64" && fieldTypeStr == "int" {
		val = val.Convert(fieldType)
	} else if fieldType != val.Type() {
		return fmt.Errorf("Provided value type " + valTypeStr + " didn't match obj field type " + fieldTypeStr)
	}
	fieldValue.Set(val)
	return nil
}

// SetStructByJSON 由json对象生成 struct
func SetStructByJSON(obj interface{}, mapData map[string]interface{}) error {
	for key, value := range mapData {
		if err := setField(obj, key, value); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func XmlToMap(r io.Reader) ParamsMap {

	params := make(ParamsMap)

	decoder := xml.NewDecoder(r)

	var (
		key   string
		value string
	)

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement: // 开始标签
			key = token.Name.Local
		case xml.CharData: // 标签内容
			content := string([]byte(token))
			value = content
		}
		if key != "xml" {
			if value != "\n" {
				params.SetString(key, value)
			}
		}
	}

	return params
}

type ParamsMap map[string]string

// map本来已经是引用类型了，所以不需要 *Params
func (p ParamsMap) SetString(k, s string) ParamsMap {
	p[k] = s
	return p
}

func (p ParamsMap) GetString(k string) string {
	s, _ := p[k]
	return s
}

func (p ParamsMap) SetInt64(k string, i int64) ParamsMap {
	p[k] = strconv.FormatInt(i, 10)
	return p
}

func (p ParamsMap) GetInt64(k string) int64 {
	i, _ := strconv.ParseInt(p.GetString(k), 10, 64)
	return i
}

// 判断key是否存在
func (p ParamsMap) ContainsKey(key string) bool {
	_, ok := p[key]
	return ok
}

// json拷贝
// param &dst新对象
// param src原对象
func JsonCopy(dst interface{}, src interface{})error  {
	marshal, err := Marshal(src)
	if err != nil{
		return err
	}
	err = Unmarshal(marshal, dst)
	if err != nil{
		return err
	}
	return nil
}