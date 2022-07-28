package py3

import "github.com/aadog/py3-go/cpy3"

var _UserException *PyObject

func UserException() *PyObject {
	return _UserException
}

func PyErr_NewException(name string,base *PyObject,dict *PyObject)*PyObject{
	return PyObjectFromInst(cpy3.PyErr_NewException(name,base.instance,dict.instance))
}

func PyExc_Exception() *PyObject {
	return PyObjectFromInst(cpy3.PyExc_Exception())
}
func PyExc_ValueError() *PyObject {
	return PyObjectFromInst(cpy3.PyExc_ValueError())
}

func PyErr_SetString(tp *PyObject, message string) {
	cpy3.PyErr_SetString(tp.instance, message)
}
