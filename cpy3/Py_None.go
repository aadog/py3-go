package cpy3

func Py_None() uintptr {
	return py_NoneStruct.Addr()
}
