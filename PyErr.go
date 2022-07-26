package py3

import "github.com/aadog/py3-go/cpy3"

func PyExc_Exception() *PyObject {
	return PyObjectFromInst(cpy3.PyExc_Exception())
}
func PyExc_ValueError() *PyObject {
	return PyObjectFromInst(cpy3.PyExc_ValueError())
}

func PyErr_SetString(tp *PyObject, message string) {
	cpy3.PyErr_SetString(tp.instance, message)
}
