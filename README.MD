
### go python bind

#### cgo-less 

```
go get github.com/aadog/py3-go
```


### test

```

func main(){
	PyImport_AppendInittab("_test", func() *PyObject {
		m := CreateModule("_test", "aa")
		m.AddFunction("add", func(a, b int) int {
			return a + b
		})
		return m.AsObj()
	})
	cpy3.Py_SetProgramName(os.Args[0])
	cpy3.Py_SetPythonHome("./")
	cpy3.Py_Initialize()
	cpy3.PyRun_SimpleString(`
        import _test
        print(_test.Call('add',1,2))
    `)
}
```