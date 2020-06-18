// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package base64

import "testing"

func TestB64EnCode(t *testing.T) {
	code := "hahaha"
	encode := B64EnCode(code)
	decode, err := B64Decode(encode)
	if err != nil {
		t.Errorf("decode err: %v", err)
	}
	if code != decode {
		t.Errorf("encode err: %v --> %v", code, encode)
	}
}

func TestB32EnCode(t *testing.T) {
	code := "hahaha"
	encode := B32EnCode(code)
	decode, err := B32Decode(encode)
	if err != nil {
		t.Errorf("decode err: %v", err)
	}
	if code != decode {
		t.Errorf("encode err: %v --> %v", code, encode)
	}
}
