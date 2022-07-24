package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type PyLong struct {
	PyObject
}

// PyLongFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyLongFromInst(inst uintptr) *PyLong {
	dl := new(PyLong)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
func PyLongFromObj(obj *PyObject) *PyLong {
	dl := new(PyLong)
	dl.PyObject = *obj
	return dl
}
func PyLongFromLong(n int) *PyLong {
	return PyLongFromInst(cpy3.PyLong_FromLong(n))
}
func PyLongFromLongLong(n int64) *PyLong {
	return PyLongFromInst(cpy3.PyLong_FromLongLong(n))
}
func PyLong_FromDouble(n float64) *PyLong {
	return PyLongFromInst(cpy3.PyLong_FromDouble(n))
}
