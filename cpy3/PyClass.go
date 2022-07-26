package cpy3

func PyType_Type() uintptr {
	return pyType_Type.Addr()
}

func PyObject_CallFunctionObjArgs(callable ...uintptr) uintptr {
	r, _, _ := pyObject_CallFunctionObjArgs.Call(callable...)
	return r
}
