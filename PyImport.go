package py3

import "github.com/aadog/py3-go/cpy3"

func PyImport_Import(name string) *PyModule {
	pName := PyUnicode_DecodeFSDefault(name)
	defer pName.DecRef()
	return PyModuleFromInst(cpy3.PyImport_Import(pName.instance))
}

func PyImport_AppendInittab(name string, initFunc func() *PyObject) int {
	r := cpy3.PyImport_AppendInittab(name, NewModuleInitFuncCallBack(name, func() *PyObject {
		obj := initFunc()
		return obj
	}))
	return r
}
