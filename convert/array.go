// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package convert

// StringArrayContains 字符串数组是否包含某字符串
func StringArrayContains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
