// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package convert

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

// Str2Int string to int
func Str2Int(s string) int {
	si, _ := strconv.Atoi(s)
	return si
}

// Str2Int32 string to int32
func Str2Int32(s string) int32 {
	si32, _ := strconv.ParseInt(s, 10, 32)
	return int32(si32)
}

// Str2Int64 string to int64
func Str2Int64(s string) int64 {
	si64, _ := strconv.ParseInt(s, 10, 64)
	return si64
}

// Str2Float64 string to float64
func Str2Float64(s string) float64 {
	sf, _ := strconv.ParseFloat(s, 64)
	return sf
}

// Str2Byte string to byte
func Str2Byte(s string) []byte {
	return []byte(s)
}

// Int642Str int64 to string
func Int642Str(i int64) string {
	return strconv.FormatInt(i, 10)
}

// InfToInt ...
func InfToInt(inter interface{}) (i int) {
	switch inter.(type) {
	case int:
		i = inter.(int)
		break
	}
	return
}

// Struct2Map ...
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// Struct2Json2Map ...
func Struct2Json2Map(obj interface{}) (result map[string]interface{}, err error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonBytes, &result)
	return
}

// Map2String ...
func Map2String(data []string) (result string) {
	if len(data) <= 0 {
		return
	}
	for _, v := range data {
		if strings.Contains(v, "\"") {
			result += v
		} else {
			result += "\"" + v + "\""
		}
		result += ","
	}
	result = strings.Trim(result, ",")
	return
}
