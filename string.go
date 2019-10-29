// Package utils 常用的方法集合
// 作者: 无双(204966@qq.com)
package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// StringJoin 高性能拼接字符串方法
func StringJoin(vars ...string) string {
	if len(vars) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	var str string
	for _, str = range vars {
		buffer.WriteString(str)
	}
	return buffer.String()
}

// StringRepeat 重复字符串
func StringRepeat(s string, count int64) string {
	strCount := strconv.FormatInt(count, 10)
	intCount, err := strconv.Atoi(strCount)
	if err != nil {
		intCount = 1
	}
	return strings.Repeat(s, intCount)
}

// Random 随机字符串
func Random(length int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandomNumber 随机数字字符串
func RandomNumber(length int64) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandomPassword 随机密码生成器
func RandomPassword(length int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!#$%^&*()"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// Substr 类似php的substr方法，截取字符串
func Substr(s string, vars ...int) string {
	var start, length int
	var end int
	bt := []rune(s)
	if len(vars) == 2 {
		start = vars[0]
		if start < 0 {
			start = len(bt) + start
			if start < 0 {
				start = 0
			}
		}
		length = vars[1]
		if length < 0 {
			end = len(bt) + length - 1
			if end < start {
				end = start
			}
			return string(bt[start:end])
		}
		if length+start > len(bt) {
			return string(bt[start:])
		}
		end = start + length
		return string(bt[start:end])
	}
	start = vars[0]
	if start < 0 {
		start = len(bt) + start
		if start < 0 {
			start = 0
		}
		return string(bt[start:len(bt)])
	}
	if start > len(bt) {
		return string(bt[len(bt):len(bt)])
	}
	return string(bt[start:])
}

// Md5 md5字符串生成
func Md5(password string) string {
	mx := md5.New()
	mx.Write([]byte(password))
	return hex.EncodeToString(mx.Sum(nil))
}

// SuperHash 字符串转hash值 uint64
func SuperHash(str string) uint64 {
	return uint64(HashID(str))<<32 + uint64(HashKey(str))
}

// HashID 计算字符串的HashID
func HashID(str string) uint32 {
	s := strings.Split(str, ".")
	var hashnum uint32
	var buf []byte
	for _, v := range s {
		if len(v) > 0 {
			buf = append(buf, byte(len(v)))
			buf = append(buf, []byte(v)...)
		}
	}
	for _, v := range buf {
		hashnum = hashnum*31 + uint32(v)
	}
	return hashnum & 0x7fffffff
}

// HashKey 计算字符串的HashKey
func HashKey(str string) uint32 {
	s := strings.Split(str, ".")
	var hashnum uint32 = 5381
	var buf []byte
	for _, v := range s {
		if len(v) > 0 {
			buf = append(buf, byte(len(v)))
			buf = append(buf, []byte(v)...)
		}
	}
	for _, v := range buf {
		hashnum += (hashnum << 5) + uint32(v)
	}
	return hashnum & 0x7fffffff
}

// IsNumbers 方法 判断是否符合这类格式的数字串, 例 123 或 123,234 符合返回true 不符合返回 false
func IsNumbers(str string) bool {
	var reg = regexp.MustCompile(`^[0-9,]+$`)
	var regNumber = regexp.MustCompile(`^[0-9]+$`)
	if reg.Match([]byte(str)) == true {
		tmp := strings.Split(str, ",")
		for _, v := range tmp {
			if regNumber.Match([]byte(v)) != true {
				return false
			}
		}
		return true
	}
	return false
}
