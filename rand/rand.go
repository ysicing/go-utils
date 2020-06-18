// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package rand

import "math/rand"

func Rand() int {
	return rand.Int()
}

func RandByNum(num int) int {
	return rand.Intn(num)
}
