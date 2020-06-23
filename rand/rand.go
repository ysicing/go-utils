// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package rand

import (
	"math/rand"
	"time"
)

// // Rand 随机数
func Rand() int {
	return rand.Int()
}

// RandByNum 随机数
func RandByNum(num int) int {
	return rand.Intn(num)
}

// RandString 生成随机字符串
func RandString(len int) string {
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
