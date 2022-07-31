package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"path/filepath"
)

func RunAnyFile(pathstr string) int {
	pathobj := NewPyUnicode(pathstr)
	defer pathobj.DecRef()
	fp := cpy3.Py_fopen_obj(pathobj.instance, "r")
	return cpy3.PyRun_AnyFile(fp, filepath.Base(pathstr))
}
func RunSimpleFile(pathstr string) int {
	pathobj := NewPyUnicode(pathstr)
	defer pathobj.DecRef()
	fp := cpy3.Py_fopen_obj(pathobj.instance, "r")
	return cpy3.PyRun_SimpleFile(fp, filepath.Base(pathstr))
}
func RunSimpleString(command string) int {
	return cpy3.PyRun_SimpleString(command)
}
