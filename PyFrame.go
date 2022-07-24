package py3_go

import (
	"github.com/aadog/go-py3/cpy3"
	"unsafe"
)

type PyFrame struct {
	PyObject
}

func (p *PyFrame) GetBack() *PyObject {
	return PyObjectFromInst(cpy3.PyFrame_GetBack(p.instance))
}
func (p *PyFrame) GetCode() *PyObject {
	return PyObjectFromInst(cpy3.PyFrame_GetCode(p.instance))
}

// PyFrameFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyFrameFromInst(inst uintptr) *PyFrame {
	dl := new(PyFrame)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
