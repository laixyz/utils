package utils

import (
	"fmt"
	"strconv"
)

// Byte2string 转换[]uint8 成 string类型
func Byte2string(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

// StringToInt64 string to int64
func StringToInt64(str string) (i int64, err error) {
	i, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// StringToUInt64 string to uint64
func StringToUInt64(str string) (i uint64, err error) {
	i, err = strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// StringToInt32 string to int32
func StringToInt32(str string) (result int32, err error) {
	var newNumber int64
	newNumber, err = strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	result = int32(newNumber)
	return result, nil
}

// StringToUint32 string to uint32
func StringToUint32(str string) (result uint32, err error) {
	var newNumber uint64
	newNumber, err = strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	result = uint32(newNumber)
	return result, nil
}

// Int64ToUint32 int64 to uint32
func Int64ToUint32(number int64) (result uint32, err error) {
	var newNumber uint64
	newNumber, err = strconv.ParseUint(fmt.Sprintf("%d", number), 10, 32)
	if err != nil {
		return 0, err
	}
	result = uint32(newNumber)
	return result, nil
}

// Int64ToInt32 int64 to int32
func Int64ToInt32(number int64) (result int32, err error) {
	var newNumber int64
	newNumber, err = strconv.ParseInt(fmt.Sprintf("%d", number), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(newNumber), nil
}
