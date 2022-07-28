
### go python3 bind
* ```python version:3.10.5```


### support list
- `cogo-less,It does not need CGO support, so the compilation speed is very fast`
- `python3 all plugin,include Native module`
- `Automatic type conversion Go to PyObject`
- `Automatic type conversion PyObject to Go`
- `import Any Python package`
- `C function is registered to python, painless interaction`

### cgo-less 

### install

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


### exception

```

func main(){
	PyImport_AppendInittab("_test", func() *PyObject {
		m := CreateModule("_test", "aa")
		m.AddFunction("add", func(a, b int) int {
		    py3.PyErr_SetString(py3.UserException(),"raise exception")
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