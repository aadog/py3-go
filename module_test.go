package py3

import (
	"fmt"
	"testing"
)

func TestCreateModule(t *testing.T) {

	//	moduleName := "test"
	//	Import_AppendInittab(moduleName, func() *PyObject {
	//		def := NewModuleDef(moduleName, "")
	//		def.AddMethodDef(NewMethodDef("aa", func() {
	//			fmt.Println("aa")
	//		}, "", 1))
	//		return CreateModule(def)
	//	})
	//	cpy3.Py_Initialize()
	//	cpy3.PyRun_SimpleString(`
	//import test
	//`)
}

type K struct {
	Doc string
}

func (v *K) Test(self *PyObject, args *PyObject) *PyObject {
	return nil
}
func TestRegModule(t *testing.T) {
	PyImport_AppendInittab("test", func() *PyObject {
		return RegPyModule("test", &K{})
	})
	fmt.Println()
}
