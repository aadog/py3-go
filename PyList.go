package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type PyList struct {
	PyObject
}

func (p *PyList) SetItem(index int64, o *PyObject) int {
	return cpy3.PyList_SetItem(p.instance, index, o.instance)
}
func (p *PyList) GetItem(index int64) *PyObject {
	return PyObjectFromInst(cpy3.PyList_GetItem(p.instance, index))
}
func (p *PyList) GetSlice(low int64, high int64) *PyObject {
	return PyObjectFromInst(cpy3.PyList_GetSlice(p.instance, low, high))
}
func (p *PyList) Size() int64 {
	return cpy3.PyList_Size(p.instance)
}
func (p *PyList) Insert(index int64, item *PyObject) int {
	return cpy3.PyList_Insert(p.instance, index, item.instance)
}
func (p *PyList) Append(item *PyObject) int {
	return cpy3.PyList_Append(p.instance, item.instance)
}

// PyListFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyListFromInst(inst uintptr) *PyList {
	dl := new(PyList)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
func PyListFromObj(obj *PyObject) *PyList {
	dl := new(PyList)
	dl.PyObject = *obj
	return dl
}
func NewPyList(len int64) *PyList {
	return PyListFromInst(cpy3.PyList_New(len))
}
