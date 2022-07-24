package py3

import (
	"fmt"
	"github.com/aadog/py3-go/cpy3"
	"sync"
	"syscall"
	"unsafe"
)

var ModuleMethodsDef = sync.Map{}
var ModuleDefMap = sync.Map{}
var PyMethodMap = sync.Map{}
var pyModuleInitMap = sync.Map{}

var PyMethodCallBack = syscall.NewCallback(func(self uintptr, args uintptr) uintptr {
	pyArgs := PyTupleFromInst(args)
	pyArgsLen := pyArgs.Size()
	if pyArgsLen < 1 {
		return Py_RETURN_NONE().instance
	}
	code := pyArgs.GetItem(0).Str()

	ifn, ok := PyMethodMap.Load(code)
	if ok == false {
		return Py_RETURN_NONE().Instance()
	}
	//处理参数
	newArgs := pyArgs.GetSlice(1, pyArgsLen)
	defer newArgs.DecRef()
	fn := ifn.(func(self *PyObject, args *PyObject) *PyObject)
	return fn(PyObjectFromInst(self), PyObjectFromInst(newArgs.Instance())).Instance()
})

type PyMethodType = func(self *PyObject, args *PyObject) *PyObject

type PyModule struct {
	PyObject
}

func (p *PyModule) GetName() string {
	return cpy3.PyModule_GetName(p.instance)
}
func (p *PyModule) GetDict() *PyDict {
	return PyDictFromInst(cpy3.PyModule_GetDict(p.instance))
}
func (p *PyModule) AddFunctions(functionsDef []PyMethodDef) int {
	methods := make([]cpy3.PyMethodDef, 0)
	moduleName := p.GetName()
	for _, method := range functionsDef {
		methodName := method.Name
		methods = append(methods, cpy3.PyMethodDef{
			Ml_name:  cpy3.GoStrToCStr(method.Name),
			Ml_meth:  NewMethodCallBack(moduleName, methodName, method.Method),
			Ml_flags: method.flags,
			Ml_doc:   cpy3.GoStrToCStr(method.Doc),
		})
	}
	methods = append(methods, cpy3.PyMethodDef{
		Ml_name:  0,
		Ml_meth:  0,
		Ml_flags: 0,
		Ml_doc:   0,
	})
	return cpy3.PyModule_AddFunctions(p.instance, uintptr(unsafe.Pointer(&methods[0])))
}
func (p *PyModule) AddIntConstant(name string, value int64) int {
	return cpy3.PyModule_AddIntConstant(p.instance, name, value)
}
func (p *PyModule) AddStringConstant(name string, value string) int {
	return cpy3.PyModule_AddStringConstant(p.instance, name, value)
}
func (p *PyModule) AddObject(name string, value *PyObject) int {
	return cpy3.PyModule_AddObject(p.instance, name, value.instance)
}
func (p *PyModule) AddObjectRef(name string, value *PyObject) int {
	return cpy3.PyModule_AddObjectRef(p.instance, name, value.instance)
}

// PyModuleFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyModuleFromInst(inst uintptr) *PyModule {
	dl := new(PyModule)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}

func NewModuleInitFuncCallBack(moduleName string, fn func() *PyObject) uintptr {
	c := syscall.NewCallback(func() uintptr {
		return fn().Instance()
	})
	pyModuleInitMap.Store(moduleName, c)
	return c
}
func NewMethodCallBack(moduleName string, methodName string, fn func(self *PyObject, args *PyObject) *PyObject) uintptr {
	PyMethodMap.Store(fmt.Sprintf("%s.%s", moduleName, methodName), fn)
	return PyMethodCallBack
}

func CreateModule(def *PyModuleDef) *PyObject {
	moduleName := def.Name
	methods := make([]cpy3.PyMethodDef, 0)
	for _, method := range def.MethodDefs {
		methodName := method.Name
		NewMethodCallBack(moduleName,methodName,method.Method)
		//methodDef := cpy3.PyMethodDef{
		//	Ml_name:  cpy3.GoStrToCStr(method.Name),
		//	Ml_meth:  NewMethodCallBack(moduleName, methodName, method.Method),
		//	Ml_flags: method.flags,
		//	Ml_doc:   cpy3.GoStrToCStr(method.Doc),
		//}
		//methods = append(methods, methodDef)
	}
	methodCallDef := cpy3.PyMethodDef{
		Ml_name:  cpy3.GoStrToCStr("Call"),
		Ml_meth:  PyMethodCallBack,
		Ml_flags: 1,
		Ml_doc:   cpy3.GoStrToCStr("跳转程序"),
	}
	methods = append(methods, methodCallDef)
	moduleNullMethodDef := cpy3.PyMethodDef{
		Ml_name:  0,
		Ml_meth:  0,
		Ml_flags: 0,
		Ml_doc:   0,
	}
	methods = append(methods, moduleNullMethodDef)

	ModuleMethodsDef.Store(fmt.Sprintf("%s", moduleName), methods)
	moduleDef := cpy3.PyModuleDef{
		M_base: cpy3.PyModuleDef_Base{
			Ob_base: cpy3.PyObject_HEAD_INIT(0),
		},
		M_name:     cpy3.GoStrToCStr(def.Name),
		M_doc:      cpy3.GoStrToCStr(def.Doc),
		M_size:     -1,
		M_methods:  uintptr(unsafe.Pointer(&methods[0])),
		M_slots:    0,
		M_traverse: 0,
		M_clear:    0,
		M_free:     0,
	}
	ModuleDefMap.Store(moduleName, def)

	pyObj := PyObjectFromInst(cpy3.PyModule_Create2(uintptr(unsafe.Pointer(&moduleDef)), 1013))
	//不知道为啥非要IncRef一次才不会崩溃
	pyObj.IncRef()
	return pyObj
}
