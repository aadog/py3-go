package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type PyNone struct {
	PyObject
}

// PyNoneFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyNoneFromInst(inst uintptr) *PyNone {
	dl := new(PyNone)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}

func Py_RETURN_NONE() *PyNone {
	none := PyNoneFromInst(cpy3.Py_None())
	none.IncRef()
	return none
}
