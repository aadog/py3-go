
### go python bind

#### cgo-less 

```
go get github.com/aadog/py3-go
```


### test

```
type K struct {
	Doc string
}

func (v *K) Test(self *PyObject, args *PyObject) *PyObject {
	fmt.Println("test")
	return nil
}
func main(){
	PyImport_AppendInittab("_test", func() *PyObject {
		return RegPyModule("_test", &K{})
	})
	cpy3.Py_SetProgramName(os.Args[0])
	cpy3.Py_SetPythonHome("./")
	cpy3.Py_Initialize()
	cpy3.PyRun_SimpleString(`
import _test
print(_test.Call('Test'))
    `)
}
```