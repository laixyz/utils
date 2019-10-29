package utils

import (
	jsoniter "github.com/json-iterator/go"
)

// JSONEncode 方法 对像转换成json 字符串格式
func JSONEncode(p interface{}) (str string, err error) {
	var b []byte
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err = json.Marshal(&p)
	return string(b), err
}

// JSONDecode 方法 json字符串转对像
func JSONDecode(str string, p interface{}) (err error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal([]byte(str), p)
}

// ToJSONString 对像 JSONEncode方法的别名
func ToJSONString(data interface{}) (string, error) {
	return JSONEncode(data)
}
