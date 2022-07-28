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
	py3.Initialize()
	cpy3.PyRun_SimpleString(PyMain)
	py3.Finalize()
}
