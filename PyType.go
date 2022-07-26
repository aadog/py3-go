package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type PyType struct {
	PyObject
}

func (p *PyType) Name() string {
	return cpy3.PyType_Name(p.instance)
}

func (p *PyType) GetModule() *PyModule {
	return PyModuleFromInst(cpy3.PyType_GetModule(p.instance))
}

// PyTypeFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyTypeFromInst(inst uintptr) *PyType {
	dl := new(PyType)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
func PyTypeFromObj(obj *PyObject) *PyType {
	dl := new(PyType)
	dl.PyObject = *obj
	return dl
}
