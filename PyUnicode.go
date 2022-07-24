package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"unsafe"
)

type PyUnicode struct {
	PyObject
}

func (p *PyUnicode) GetLength() int64 {
	return cpy3.PyUnicode_GetLength(p.instance)
}

func (p *PyObject) AsUTF8() string {
	return cpy3.PyUnicode_AsUTF8(p.instance)
}

func PyUnicode_DecodeFSDefault(s string) *PyUnicode {
	return PyUnicodeFromInst(cpy3.PyUnicode_DecodeFSDefault(s))
}
func PyUnicode_FromString(s string) *PyUnicode {
	return PyUnicodeFromInst(cpy3.PyUnicode_FromString(s))
}

// PyUnicodeFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyUnicodeFromInst(inst uintptr) *PyUnicode {
	dl := new(PyUnicode)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
