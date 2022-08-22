package cpy3

import (
	"github.com/aadog/msvcrt-go"
	"unsafe"
)

func PyUnicode_DecodeFSDefault(u string) uintptr {
	r, _, _ := pyUnicode_DecodeFSDefault.Call(msvcrt.StringToCUTF8String(u))
	return r
}
func PyUnicode_FromString(u string) uintptr {
	if u == "" {
		b := make([]byte, 1)
		b[0] = 0x0
		r, _, _ := pyUnicode_FromString.Call(uintptr(unsafe.Pointer(&b[0])))
		return r
	}
	r, _, _ := pyUnicode_FromString.Call(msvcrt.StringToCUTF8String(u))
	return r
}
func PyUnicode_GetLength(obj uintptr) int64 {
	r, _, _ := pyUnicode_GetLength.Call(obj)
	return int64(r)
}
func PyUnicode_AsUTF8(obj uintptr) string {
	r, _, _ := pyUnicode_AsUTF8.Call(obj)
	return msvcrt.CUtf8ToString(r)
}
