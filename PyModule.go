package py3

import (
	"fmt"
	"github.com/aadog/py3-go/cpy3"
	"reflect"
	"sync"
	"syscall"
	"unsafe"
)

const (
	PYTHON_API_VERSION = 1013
	PYTHON_API_STRING  = "1013"
)

var pyModuleInitMap = sync.Map{}

type PyModuleGoObj struct {
	moduleDef *cpy3.PyModuleDef
	CallMap   sync.Map
}
type PyModule struct {
	PyObject
	GoObj *PyModuleGoObj
}

func (p *PyModule) GetName() string {
	return cpy3.PyModule_GetName(p.instance)
}
func (p *PyModule) GetDict() *PyDict {
	return PyDictFromInst(cpy3.PyModule_GetDict(p.instance))
}

func (p *PyModule) AddFunction(name string, fn interface{}) {
	p.GoObj.CallMap.Store(name, fn)
}
func (p *PyModule) AddClass(class *PyClass) {
	className := class.GetAttrString("__className__")
	p.AddObject(className.AsUTF8(), class.AsObj())
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

	name := dl.GetName()
	smmodule, _ := SystemModuleMap.Load(name)
	dl.GoObj = smmodule.(*PyModule).GoObj
	return dl
}
func PyModuleFromObj(obj *PyObject) *PyModule {
	dl := new(PyModule)
	dl.PyObject = *obj

	name := dl.GetName()
	smmodule, _ := SystemModuleMap.Load(name)
	dl.GoObj = smmodule.(*PyModule).GoObj
	return dl
}

func PyTypeToGoType(p *PyObject) any {

	return 0
}
func PyMethodForward(self *PyModule, args *PyTuple, method interface{}) *PyObject {
	methodType := reflect.TypeOf(method)
	methodValue := reflect.ValueOf(method)
	if methodType.Kind() != reflect.Func {
		return Py_RETURN_NONE().AsObj()
	}
	if int64(methodType.NumIn()) != args.Size() {
		PyErr_SetString(UserException(), fmt.Sprintf("The number of parameters does not match,%d parameter is required, and you have entered %d", methodType.NumIn(), args.Size()))
		return Py_RETURN_NONE().AsObj()
	}

	fnArgs := make([]reflect.Value, 0)
	for i := 0; i < methodType.NumIn(); i++ {
		fnArgType := methodType.In(i)
		inArg := args.GetItem(int64(i))
		fnArgs = append(fnArgs, reflect.ValueOf(PyObjectToGo(inArg, fnArgType)))
	}
	rets := methodValue.Call(fnArgs)
	if len(rets) == 1 {
		firstRet := rets[0]
		r := GoToPyObject(firstRet.Interface())
		return r
	} else if len(rets) > 1 {
		l := NewPyList(0)
		for _, r := range rets {
			obj := GoToPyObject(r.Interface())
			defer obj.DecRef()
			l.Append(obj)
		}
		return l.AsObj()
	}
	return Py_RETURN_NONE().AsObj()
}

var PyModuleMethodForwardCallBack = syscall.NewCallback(func(self uintptr, args uintptr) uintptr {

	pyArgs := PyTupleFromInst(args)
	pyArgsLen := pyArgs.Size()
	if pyArgsLen < 1 {
		return Py_RETURN_NONE().instance
	}
	ForwardCode := pyArgs.GetItem(0).Str()

	pyModule := PyModuleFromInst(self)
	moduleName := pyModule.GetName()
	ifn, ok := pyModule.GoObj.CallMap.Load(ForwardCode)
	if ok == false {
		PyErr_SetString(UserException(), fmt.Sprintf("%s not find method %s ", moduleName, ForwardCode))
		return Py_RETURN_NONE().Instance()
	}
	//处理参数
	newArgs := PyTupleFromObj(pyArgs.GetSlice(1, pyArgsLen))
	defer newArgs.DecRef()

	return PyMethodForward(pyModule, newArgs, ifn).Instance()
})

func CreateModule(name string, doc string) *PyModule {
	var pyModuleMethodCallDefs = make([]cpy3.PyMethodDef, 0)
	methodCallDef := cpy3.PyMethodDef{
		Ml_name:  cpy3.GoStrToCStr("Call"),
		Ml_meth:  PyModuleMethodForwardCallBack,
		Ml_flags: 1,
		Ml_doc:   cpy3.GoStrToCStr("module call forward"),
	}
	pyModuleMethodCallDefs = append(pyModuleMethodCallDefs, methodCallDef)
	moduleNullMethodDef := cpy3.PyMethodDef{
		Ml_name:  0,
		Ml_meth:  0,
		Ml_flags: 0,
		Ml_doc:   0,
	}
	pyModuleMethodCallDefs = append(pyModuleMethodCallDefs, moduleNullMethodDef)

	module := &PyModule{}
	module.GoObj = new(PyModuleGoObj)
	module.GoObj.moduleDef = &cpy3.PyModuleDef{
		M_base: cpy3.PyModuleDef_Base{
			Ob_base: cpy3.PyObject_HEAD_INIT(0),
		},
		M_name:     cpy3.GoStrToCStr(name),
		M_doc:      cpy3.GoStrToCStr(doc),
		M_size:     -1,
		M_methods:  uintptr(unsafe.Pointer(&pyModuleMethodCallDefs[0])),
		M_slots:    0,
		M_traverse: 0,
		M_clear:    0,
		M_free:     0,
	}
	ptr := cpy3.PyModule_Create2(uintptr(unsafe.Pointer(module.GoObj.moduleDef)), PYTHON_API_VERSION)
	module.instance = ptr
	module.ptr = unsafe.Pointer(module.instance)
	SystemModuleMap.Store(name, module)
	return module
}

func NewModuleInitFuncCallBack(moduleName string, fn func() *PyObject) uintptr {
	var c = syscall.NewCallback(func() uintptr {
		return fn().Instance()
	})
	pyModuleInitMap.Store(moduleName, c)
	return c
}
