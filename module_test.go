package py3

import (
	"fmt"
	"github.com/aadog/py3-go/cpy3"
	"os"
	"testing"
)

type K struct {
	Doc string
}

func (v *K) Test(a int, b int) int {
	fmt.Println("test")
	return a + b
}

func TestGoModule(t *testing.T) {

	SetProgramName(os.Args[0])
	SetPythonHome("./")
	InitGoModuleName("go", "go")
	Initialize()
	defer Finalize()
	fmt.Println(SystemModuleMap)
	RunSimpleString(`
import go
print(go)
`)
}

func TestRegModule(t *testing.T) {

	PyImport_AppendInittab("_test", func() *PyObject {
		m := CreateModule("_test", "aa")
		m.AddFunction("add", func(a, b int) int {
			return a + b
		})
		return m.AsObj()
	})
	SetProgramName(os.Args[0])
	SetPythonHome("./")
	Initialize()
	defer Finalize()
	RunSimpleString(`
import _test
print(_test.Call('add',1,2))
`)
}

func TestRunSimpleFile(t *testing.T) {
	cpy3.Py_SetProgramName(os.Args[0])
	cpy3.Py_SetPythonHome("./")
	Initialize()
	RunSimpleFile("./testdata/RunSimpleFile/test.py")
}

func TestRunAnyFile(t *testing.T) {
	cpy3.Py_SetProgramName(os.Args[0])
	cpy3.Py_SetPythonHome("./")
	Initialize()
	RunAnyFile("./testdata/RunAnyFile/test.py")
}

func TestGoStruct(t *testing.T) {

	SetProgramName(os.Args[0])
	SetPythonHome("./")
	InitGoModuleName("go", "go")
	Initialize()
	defer Finalize()
	RunSimpleString(`
import go
print(go.MyClass().Call("z","22"))
#x=go.MyClass()
#print(x.Call("mb","xx"))
`)
}
