package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type IPyObject interface {
	SetObject(object *PyObject)
}

type PyObject struct {
	CObj
}

func (p *PyObject) String() string {
	return p.Str()
}
func (p *PyObject) GetAttr(attr_name string) *PyObject {
	name := PyUnicode_FromString(attr_name)
	defer name.DecRef()
	o := cpy3.PyObject_GetAttr(p.instance, name.instance)
	return PyObjectFromInst(o)
}
func (p *PyObject) Str() string {
	o := PyUnicodeFromInst(cpy3.PyObject_Str(p.instance))
	defer o.DecRef()
	return cpy3.PyUnicode_AsUTF8(o.instance)
}
func (o *PyObject) DecRef() {
	cpy3.Py_DecRef(o.instance)
}
func (o *PyObject) IncRef() {
	if o.instance != 0 {
		cpy3.Py_IncRef(o.instance)
	}
}
func (p *PyObject) AsObj() *PyObject {
	return p
}

// PyObjectFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyObjectFromInst(inst uintptr) *PyObject {
	dl := new(PyObject)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
