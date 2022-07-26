package cpy3

import "github.com/aadog/msvcrt-go"

func PyDict_SetItem(obj uintptr, key uintptr, val uintptr) int {
	r, _, _ := pyDict_SetItem.Call(obj, key, val)
	return int(r)
}

func PyDict_SetItemString(obj uintptr, key string, val uintptr) int {
	r, _, _ := pyDict_SetItemString.Call(obj, msvcrt.StringToCUTF8String(key), val)
	return int(r)
}

func PyDict_New() uintptr {
	r, _, _ := pyDict_New.Call()
	return r
}

func PyDict_Size(obj uintptr) int64 {
	r, _, _ := pyDict_Size.Call(obj)
	return int64(r)
}

func PyDict_Clear(obj uintptr) {
	pyDict_Clear.Call(obj)
}
func PyDict_GetItem(obj uintptr, key uintptr) uintptr {
	r, _, _ := pyDict_GetItem.Call(obj, key)
	return r
}
func PyDict_Keys(obj uintptr) uintptr {
	r, _, _ := pyDict_Keys.Call(obj)
	return r
}
func PyDict_GetItemString(obj uintptr, key string) uintptr {
	r, _, _ := pyDict_GetItemString.Call(obj, msvcrt.StringToCUTF8String(key))
	return r
}

func PyDict_DelItem(obj uintptr, key uintptr) int {
	r, _, _ := pyDict_DelItem.Call(obj, key)
	return int(r)
}
func PyDict_DelItemString(obj uintptr, key string) int {
	r, _, _ := pyDict_DelItemString.Call(obj, msvcrt.StringToCUTF8String(key))
	return int(r)
}
