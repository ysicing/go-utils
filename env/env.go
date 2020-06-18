// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package env

import (
	"os"
)

// EnvString 获取环境变量
func EnvString(env, fallback string) string {
	if e := os.Getenv(env); e != "" {
		return e
	}
	return fallback
}
