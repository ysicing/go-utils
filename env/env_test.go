// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvString(t *testing.T) {
	const expected = "foo"

	assert := assert.New(t)

	key := "FLOCKER_SET_VAR"
	os.Setenv(key, expected)
	assert.Equal(expected, EnvString(key, "~"+expected))

	key = "FLOCKER_UNSET_VAR"
	assert.Equal(expected, EnvString(key, expected))
}
