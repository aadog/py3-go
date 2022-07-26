package cpy3

func PyObject_GetAttrString(obj uintptr, attr_name string) uintptr {
	r, _, _ := pyObject_GetAttrString.Call(obj, GoStrToCStr(attr_name))
	return r
}
func PyObject_HasAttrString(obj uintptr, attr_name string) int {
	r, _, _ := pyObject_HasAttrString.Call(obj, GoStrToCStr(attr_name))
	return int(r)
}
func PyObject_SetAttrString(obj uintptr, attr_name string, v uintptr) int {
	r, _, _ := pyObject_SetAttrString.Call(obj, GoStrToCStr(attr_name), v)
	return int(r)
}

func PyObject_DelAttrString(obj uintptr, attr_name string) int {
	r, _, _ := pyObject_DelAttrString.Call(obj, GoStrToCStr(attr_name))
	return int(r)
}

func PyObject_Type(obj uintptr) uintptr {
	r, _, _ := pyObject_Type.Call(obj)
	return r
}

func PyObject_Str(obj uintptr) uintptr {
	r, _, _ := pyObject_Str.Call(obj)
	return r
}

func PyObject_GetAttr(obj uintptr, attr_name uintptr) uintptr {
	r, _, _ := pyObject_GetAttr.Call(obj, attr_name)
	return r
}

func PyObject_Call(obj uintptr, args uintptr, kwargs uintptr) uintptr {
	r, _, _ := pyObject_Call.Call(obj, args, kwargs)
	return r
}
func PyObject_CallObject(obj uintptr, args uintptr) uintptr {
	r, _, _ := pyObject_CallObject.Call(obj, args)
	return r
}
func PyObject_CallNoArgs(obj uintptr) uintptr {
	r, _, _ := pyObject_CallNoArgs.Call(obj)
	return r
}
