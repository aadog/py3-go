package cpy3

func PyBool_FromLong(l int64) uintptr {
	r, _, _ := pyBool_FromLong.Call(uintptr(l))
	return r
}
