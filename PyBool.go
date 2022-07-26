package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type PyBool struct {
	PyObject
}

// PyBoolFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyBoolFromInst(inst uintptr) *PyBool {
	dl := new(PyBool)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
func PyBoolFromObj(obj *PyObject) *PyBool {
	dl := new(PyBool)
	dl.PyObject = *obj
	return dl
}

func NewPyBoolFromLong(l int64) *PyObject {
	return PyObjectFromInst(cpy3.PyBool_FromLong(l))
}
func NewPyBool(b bool) *PyObject {
	l := 0
	if b == true {
		l = 1
	}
	return NewPyBoolFromLong(int64(l))
}
