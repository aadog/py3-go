package py3_go

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type PyTuple struct {
	PyObject
}

func (p *PyTuple) Size() int64 {
	return cpy3.PyTuple_Size(p.instance)
}
func (p *PyTuple) Check() int64 {
	return cpy3.PyTuple_Check(p.instance)
}
func (p *PyTuple) GetItem(pos int64) *PyObject {
	return PyObjectFromInst(cpy3.PyTuple_GetItem(p.instance, pos))
}
func (p *PyTuple) GetSlice(low int64, high int64) *PyObject {
	return PyObjectFromInst(cpy3.PyTuple_GetSlice(p.instance, low, high))
}

// PyTupleFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyTupleFromInst(inst uintptr) *PyTuple {
	dl := new(PyTuple)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
func PyTupleFromObj(obj *PyObject) *PyTuple {
	dl := new(PyTuple)
	dl.PyObject = *obj
	return dl
}
