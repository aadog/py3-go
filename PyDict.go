package py3_go

import "unsafe"

type PyDict struct {
	PyObject
}

// PyDictFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyDictFromInst(inst uintptr) *PyDict {
	dl := new(PyDict)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
