package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type PyDict struct {
	PyObject
}

func (p *PyDict) DelItemString(key string) int {
	return cpy3.PyDict_DelItemString(p.instance, key)
}

func (p *PyDict) DelItem(key *PyObject) int {
	return cpy3.PyDict_DelItem(p.instance, key.instance)
}

func (p *PyDict) GetItemString(key string) *PyObject {
	return PyObjectFromInst(cpy3.PyDict_GetItemString(p.instance, key))
}

func (p *PyDict) Keys() *PyObject {
	return PyObjectFromInst(cpy3.PyDict_Keys(p.instance))
}
func (p *PyDict) GetItem(key *PyObject) *PyObject {
	return PyObjectFromInst(cpy3.PyDict_GetItem(p.instance, key.instance))
}

func (p *PyDict) SetItemString(key string, val *PyObject) int {
	return cpy3.PyDict_SetItemString(p.instance, key, val.instance)
}
func (p *PyDict) SetItem(key *PyObject, val *PyObject) int {
	return cpy3.PyDict_SetItem(p.instance, key.instance, val.instance)
}

func (p *PyDict) Size() int64 {
	return cpy3.PyDict_Size(p.instance)
}
func (p *PyDict) Clear() {
	cpy3.PyDict_Clear(p.instance)
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
func PyDictFromObj(obj *PyObject) *PyDict {
	dl := new(PyDict)
	dl.PyObject = *obj
	return dl
}

func NewPyDict() *PyDict {
	return PyDictFromInst(cpy3.PyDict_New())
}
