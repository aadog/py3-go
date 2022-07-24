package cpy3

func PyObject_Str(obj uintptr) uintptr {
	r, _, _ := pyObject_Str.Call(obj)
	return r
}

func PyObject_GetAttr(obj uintptr, attr_name uintptr) uintptr {
	r, _, _ := pyObject_GetAttr.Call(obj, attr_name)
	return r
}
