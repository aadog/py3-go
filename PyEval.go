package py3

import "github.com/aadog/py3-go/cpy3"

func PyEval_GetBuiltins() *PyObject {
	return PyObjectFromInst(cpy3.PyEval_GetBuiltins())
}
func PyEval_GetLocals() *PyObject {
	return PyObjectFromInst(cpy3.PyEval_GetLocals())
}
func PyEval_GetGlobals() *PyObject {
	return PyObjectFromInst(cpy3.PyEval_GetGlobals())
}
func PyEval_GetFrame() *PyFrame {
	return PyFrameFromInst(cpy3.PyEval_GetFrame())
}
