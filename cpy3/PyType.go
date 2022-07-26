package cpy3

func PyType_Name(obj uintptr) string {
	r, _, _ := pyType_Name.Call(obj)
	return CStrToGoStr(r)
}

func PyType_GetModule(obj uintptr) uintptr {
	r, _, _ := pyType_GetModule.Call(obj)
	return r
}
