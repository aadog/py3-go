package cpy3

import "github.com/aadog/msvcrt-go"

func PyErr_NewException(name string, base uintptr, dict uintptr) uintptr {
	r, _, _ := pyErr_NewException.Call(msvcrt.StringToCUTF8String(name), base, dict)
	return r
}
func PyExc_Exception() uintptr {
	return pyExc_Exception.Addr()
}

func PyExc_ValueError() uintptr {
	return pyExc_ValueError.Addr()
}

func PyErr_SetString(tp uintptr, message string) {
	pyErr_SetString.Call(tp, msvcrt.StringToCUTF8String(message))
}
