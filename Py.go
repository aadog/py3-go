package py3

import (
	"fmt"
	"github.com/aadog/py3-go/cpy3"
	"sync"
)

var SystemModuleMap = sync.Map{}
var GoModule *PyModule

func InitGoModuleName(name string, doc string) {
	PyImport_AppendInittab(name, func() *PyObject {
		GoModule = CreateModule(name, doc)
		cls := CreateClass("MyClass", "")
		cls.AddFunction("z", func(self *PyObject) {
			fmt.Println("z")
		})
		GoModule.AddClass(cls)

		return GoModule.AsObj()
	})
}
func Initialize() {
	//init python3
	cpy3.Py_Initialize()

	//new userexception
	_UserException = PyErr_NewException("gofunction.error", PyNil, PyNil)
}

func IsInitialized() int {
	return cpy3.Py_IsInitialized()
}
func Finalize() {
	SystemModuleMap.Range(func(key, value any) bool {
		//SystemModuleMap.Delete(key)
		m := value.(*PyModule)
		m.DecRef()
		return true
	})
	_UserException.DecRef()
	cpy3.Py_Finalize()
}
func FinalizeEx() int {
	SystemModuleMap.Range(func(key, value any) bool {
		//SystemModuleMap.Delete(key)
		m := value.(*PyModule)
		m.DecRef()
		return true
	})
	_UserException.DecRef()
	return cpy3.Py_FinalizeEx()
}
func SetProgramName(name string) {
	cpy3.Py_SetProgramName(name)
}
func SetPythonHome(home string) {
	cpy3.Py_SetPythonHome(home)
}
