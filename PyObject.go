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
func (p *PyObject) CallNoArgs() *PyObject {
	return PyObjectFromInst(cpy3.PyObject_CallNoArgs(p.Instance()))
}
func (p *PyObject) Call(args *PyObject, kwargs *PyObject) *PyObject {
	return PyObjectFromInst(cpy3.PyObject_Call(p.Instance(), args.instance, kwargs.instance))
}
func (p *PyObject) CallObject(args *PyObject) *PyObject {
	return PyObjectFromInst(cpy3.PyObject_CallObject(p.Instance(), args.instance))
}
func (p *PyObject) PyObject_Call(args *PyObject, kwargs *PyObject) *PyObject {
	return PyObjectFromInst(cpy3.PyObject_Call(p.Instance(), args.instance, kwargs.instance))
}

func (p *PyObject) GetAttr(attr_name string) *PyObject {
	name := NewPyUnicode(attr_name)
	defer name.DecRef()
	o := cpy3.PyObject_GetAttr(p.instance, name.instance)
	return PyObjectFromInst(o)
}
func (p *PyObject) GetAttrString(attr_name string) *PyObject {
	o := cpy3.PyObject_GetAttrString(p.instance, attr_name)
	return PyObjectFromInst(o)
}
func (p *PyObject) HashAttrString(attr_name string) int {
	o := cpy3.PyObject_HasAttrString(p.instance, attr_name)
	return o
}
func (p *PyObject) SetAttrString(attr_name string, v *PyObject) int {
	o := cpy3.PyObject_SetAttrString(p.instance, attr_name, v.instance)
	return o
}
func (p *PyObject) DelAttrString(attr_name string) int {
	o := cpy3.PyObject_DelAttrString(p.instance, attr_name)
	return o
}

func (p *PyObject) Str() string {
	o := PyUnicodeFromInst(cpy3.PyObject_Str(p.instance))
	defer o.DecRef()
	return cpy3.PyUnicode_AsUTF8(o.instance)
}
func (o *PyObject) DecRef() {
	if o.instance != 0 {
		cpy3.Py_DecRef(o.instance)
	}
}
func (o *PyObject) IncRef() {
	if o.instance != 0 {
		cpy3.Py_IncRef(o.instance)
	}
}
func (p *PyObject) RefCount() int {
	return int(cpy3.PyObjectFromPtr(p.instance).Ob_refcnt)
}
func (p *PyObject) AsObj() *PyObject {
	return p
}
func (p *PyObject) Type() *PyType {
	return PyTypeFromInst(cpy3.PyObject_Type(p.instance))
}
func (p *PyObject) AsInt() int {
	return int(p.AsLong())
}
func (p *PyObject) AsDouble() float64 {
	return cpy3.PyLong_AsDouble(p.instance)
}
func (p *PyObject) AsLong() int {
	return cpy3.PyLong_AsLong(p.instance)
}
func (p *PyObject) AsLongLong() int64 {
	return cpy3.PyLong_AsLongLong(p.instance)
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

var PyNil = PyObjectFromInst(0)
