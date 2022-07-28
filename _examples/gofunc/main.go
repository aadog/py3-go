package main

import (
	_ "embed"
	"github.com/aadog/py3-go"
	"github.com/aadog/py3-go/cpy3"
	"os"
)

//go:embed main.py
var PyMain string

func main() {
	cpy3.Py_SetProgramName(os.Args[0])
	cpy3.Py_SetPythonHome("./")
	py3.PyImport_AppendInittab("gofunc", func() *py3.PyObject {
		m := py3.CreateModule("gofunc", "gofunc")
		m.AddFunction("add", func(a int, b int) int {
			return a + b
		})
		m.AddFunction("py", func(a *py3.PyObject) *py3.PyObject {
			return a
		})
		return m.AsObj()
	})
	py3.Initialize()
	cpy3.PyRun_SimpleString(PyMain)
	py3.Finalize()
}
