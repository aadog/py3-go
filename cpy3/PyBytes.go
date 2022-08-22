package cpy3

import "github.com/aadog/msvcrt-go"

func PyBytes_FromString(s string) uintptr {
	r, _, _ := pyBytes_FromString.Call(msvcrt.StringToCUTF8String(s))
	return r
}
