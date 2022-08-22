package cpy3

import "github.com/aadog/msvcrt-go"

func PyType_Name(obj uintptr) string {
	r, _, _ := pyType_Name.Call(obj)
	return msvcrt.CUtf8ToString(r)
}

func PyType_GetModule(obj uintptr) uintptr {
	r, _, _ := pyType_GetModule.Call(obj)
	return r
}
