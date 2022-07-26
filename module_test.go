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

func (v *K) Test(self *PyObject, args *PyObject) *PyObject {
	fmt.Println("test")
	return nil
}
func TestRegModule(t *testing.T) {
	PyImport_AppendInittab("_test", func() *PyObject {
		return CreateModule("_test", "aa").AsObj()
	})
	cpy3.Py_SetProgramName(os.Args[0])
	cpy3.Py_SetPythonHome("./")
	cpy3.Py_Initialize()
	cpy3.PyRun_SimpleString(`
print("aaa")
#import _test
#print(_test.Call('Test'))
`)
}
