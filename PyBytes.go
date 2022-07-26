package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type PyBytes struct {
	PyObject
}

func PyBytes_FromString(s string) *PyBytes {
	return PyBytesFromInst(cpy3.PyBytes_FromString(s))
}

// PyBytesFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyBytesFromInst(inst uintptr) *PyBytes {
	dl := new(PyBytes)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
