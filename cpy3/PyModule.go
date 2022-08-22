package cpy3

import (
	"github.com/aadog/msvcrt-go"
	"unsafe"
)

type PyMethodDef struct {
	Ml_name  uintptr
	Ml_meth  uintptr
	Ml_flags int
	Ml_doc   uintptr
}
type PyObject struct {
	//Ob_next   uintptr
	//Ob_prev   uintptr
	Ob_refcnt uintptr
	Ob_type   uintptr
}

type PyModuleDef_Base struct {
	Ob_base PyObject
	M_init  uintptr
	M_index uintptr
	M_copy  uintptr
}
type PyModuleDef struct {
	M_base     PyModuleDef_Base
	M_name     uintptr
	M_doc      uintptr
	M_size     int64
	M_methods  uintptr
	M_slots    uintptr
	M_traverse uintptr
	M_clear    uintptr
	M_free     uintptr
}

func PyObjectFromPtr(ptr uintptr) *PyObject {
	return (*PyObject)(unsafe.Pointer(ptr))
}

func PyObject_HEAD_INIT(ob_type uintptr) PyObject {
	return PyObject{
		Ob_refcnt: 1,
		Ob_type:   0,
	}
}

func PyModule_Create2(PyModuleDef uintptr, apiver int) uintptr {
	r, _, _ := pyModule_Create2.Call(PyModuleDef, uintptr(apiver))
	return r
}
func PyImport_Import(name uintptr) uintptr {
	r, _, _ := pyImport_Import.Call(name)
	return r
}

func PyModule_GetName(obj uintptr) string {
	r, _, _ := pyModule_GetName.Call(obj)
	return msvcrt.CUtf8ToString(r)
}
func PyModule_GetDict(obj uintptr) uintptr {
	r, _, _ := pyModule_GetDict.Call(obj)
	return r
}
func PyModule_AddFunctions(obj uintptr, functionsDef uintptr) int {
	r, _, _ := pyModule_AddFunctions.Call(obj)
	return int(r)
}
func PyModule_AddIntConstant(obj uintptr, name string, value int64) int {
	r, _, _ := pyModule_AddIntConstant.Call(obj, msvcrt.StringToCUTF8String(name), uintptr(value))
	return int(r)
}
func PyModule_AddStringConstant(obj uintptr, name string, value string) int {
	r, _, _ := pyModule_AddStringConstant.Call(obj, msvcrt.StringToCUTF8String(name), msvcrt.StringToCUTF8String(value))
	return int(r)
}
func PyModule_AddObject(obj uintptr, name string, value uintptr) int {
	r, _, _ := pyModule_AddObject.Call(obj, msvcrt.StringToCUTF8String(name), value)
	return int(r)
}
func PyModule_AddObjectRef(obj uintptr, name string, value uintptr) int {
	r, _, _ := pyModule_AddObjectRef.Call(obj, msvcrt.StringToCUTF8String(name), value)
	return int(r)
}
