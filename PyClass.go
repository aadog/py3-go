package py3

import (
	"fmt"
	"github.com/aadog/py3-go/cpy3"
	"sync"
	"syscall"
	"unsafe"
)

var SystemClassMap = sync.Map{}

type PyClassGoObj struct {
	CallMap sync.Map
}
type PyClass struct {
	PyObject
	GoObj *PyClassGoObj
}

func (p *PyClass) GetName() string {
	return p.GetAttrString("__className__").AsUTF8()
}

func (p *PyClass) AddFunction(name string, fn interface{}) {
	p.GoObj.CallMap.Store(name, fn)
}

// PyClassFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyClassFromInst(inst uintptr) *PyClass {
	dl := new(PyClass)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)

	name := dl.GetName()
	smmodule, _ := SystemClassMap.Load(name)
	dl.GoObj = smmodule.(*PyClass).GoObj
	return dl
}
func PyClassFromObj(obj *PyObject) *PyClass {
	dl := new(PyClass)
	dl.PyObject = *obj

	name := dl.GetName()
	smmodule, _ := SystemClassMap.Load(name)
	dl.GoObj = smmodule.(*PyClass).GoObj
	return dl
}

func PyClassInstanceMethodForward(self *PyObject, args *PyTuple, method interface{}) *PyObject {
	fmt.Println("class call")
	//	methodType := reflect.TypeOf(method)
	//	methodValue := reflect.ValueOf(method)
	//	if methodType.Kind() != reflect.Func {
	//		return Py_RETURN_NONE().AsObj()
	//	}
	//	if int64(methodType.NumIn()) != args.Size() {
	//		PyErr_SetString(UserException(), fmt.Sprintf("The number of parameters does not match,%d parameter is required, and you have entered %d", methodType.NumIn(), args.Size()))
	//		return Py_RETURN_NONE().AsObj()
	//	}
	//
	//	fnArgs := make([]reflect.Value, 0)
	//	for i := 0; i < methodType.NumIn(); i++ {
	//		fnArgType := methodType.In(i)
	//		inArg := args.GetItem(int64(i))
	//		fnArgs = append(fnArgs, reflect.ValueOf(PyObjectToGo(inArg, fnArgType)))
	//	}
	//	rets := methodValue.Call(fnArgs)
	//	if len(rets) == 1 {
	//		firstRet := rets[0]
	//		r := GoToPyObject(firstRet.Interface())
	//		return r
	//	} else if len(rets) > 1 {
	//		l := NewPyList(0)
	//		for _, r := range rets {
	//			obj := GoToPyObject(r.Interface())
	//			defer obj.DecRef()
	//			l.Append(obj)
	//		}
	//		return l.AsObj()
	//	}
	return Py_RETURN_NONE().AsObj()
}

func CreateClass(name string, doc string) *PyClass {
	class := &PyClass{}
	class.GoObj = new(PyClassGoObj)

	dict := map[string]any{}
	pClassName := GoToPyObject(name)
	defer pClassName.DecRef()
	pClassBases := NewPyTuple(0)
	defer pClassBases.DecRef()
	pClassDic := GoToPyObject(dict)
	defer pClassDic.DecRef()
	class.instance = cpy3.PyObject_CallFunctionObjArgs(cpy3.PyType_Type(), pClassName.Instance(), pClassBases.Instance(), pClassDic.Instance(), 0)
	class.ptr = unsafe.Pointer(class.instance)
	class.SetAttrString("__className__", pClassName)

	fn := PyObjectFromInst(cpy3.PyCFunction_New(uintptr(unsafe.Pointer(PyClassInstanceMethodCallDef)), PyNil.instance))
	defer fn.DecRef()
	method := PyObjectFromInst(cpy3.PyInstanceMethod(fn.instance))
	defer method.DecRef()
	class.SetAttrString("Call", method)

	SystemClassMap.Store(name, class)
	return class
}

var PyClassInstanceMethodCallDef *cpy3.PyMethodDef
var PyClassInstanceMethodForwardCallBack = syscall.NewCallback(func(self uintptr, args uintptr) uintptr {
	pyArgs := PyTupleFromInst(args)
	arg1 := pyArgs.GetItem(0)
	tp := arg1.Type()
	defer tp.DecRef()
	if tp.Name() == "str" {
		fmt.Println("static")
	} else {
		pyArgsLen := pyArgs.Size()
		if pyArgsLen < 2 {
			return Py_RETURN_NONE().instance
		}
		ForwardCode := pyArgs.GetItem(1).Str()

		//处理参数
		newArgs := PyTupleFromObj(pyArgs.GetSlice(2, pyArgsLen))
		defer newArgs.DecRef()
		pySelf := pyArgs.GetItem(0)

		PyClass := PyClassFromObj(pySelf)
		className := PyClass.GetName()
		ifn, ok := PyClass.GoObj.CallMap.Load(ForwardCode)
		if ok == false {
			PyErr_SetString(UserException(), fmt.Sprintf("%s not find method %s ", className, ForwardCode))
			return Py_RETURN_NONE().Instance()
		}

		return PyClassInstanceMethodForward(pySelf, newArgs, ifn).instance
	}
	return Py_RETURN_NONE().instance
})

//func init() {
//	PyClassInstanceMethodCallDef = &cpy3.PyMethodDef{
//		Ml_name:  cpy3.GoStrToCStr("Call"),
//		Ml_meth:  PyClassInstanceMethodForwardCallBack,
//		Ml_flags: 3,
//		Ml_doc:   cpy3.GoStrToCStr("class call forward"),
//	}
//}

//func CreateClass(name string, dict map[string]any) *PyClass {
//	if dict == nil {
//		dict = map[string]any{}
//	}
//	pClassName := GoToPyObject(name)
//	defer pClassName.DecRef()
//	pClassBases := NewPyTuple(0)
//	defer pClassBases.DecRef()
//	pClassDic := GoToPyObject(dict)
//	defer pClassDic.DecRef()
//	pClass := PyObjectFromInst(cpy3.PyObject_CallFunctionObjArgs(cpy3.PyType_Type(), pClassName.Instance(), pClassBases.Instance(), pClassDic.Instance(), 0))
//	pClass.SetAttrString("__className__", NewPyUnicode(name).AsObj())
//
//	fnName := "Call"
//	methodCallDef := &cpy3.PyMethodDef{
//		Ml_name: cpy3.GoStrToCStr(fnName),
//		Ml_meth: syscall.NewCallback(func(self uintptr, args uintptr) uintptr {
//			pyArgs := PyTupleFromInst(args)
//			pyArgsLen := pyArgs.Size()
//			if pyArgsLen < 2 {
//				return Py_RETURN_NONE().instance
//			}
//			pySelf := pyArgs.GetItem(0)
//			ForwardCode := pyArgs.GetItem(1).Str()
//
//			ifn, ok := PyMethodMap.Load(ForwardCode)
//			if ok == false {
//				return Py_RETURN_NONE().Instance()
//			}
//			//处理参数
//			newArgs := PyTupleFromObj(pyArgs.GetSlice(1, pyArgsLen))
//			defer newArgs.DecRef()
//			return PyInstanceMethodForward(pySelf, newArgs, ifn).Instance()
//		}),
//		Ml_flags: 3,
//		Ml_doc:   cpy3.GoStrToCStr("module call forward"),
//	}
//	PyMethodMap.Store("x", methodCallDef)
//
//	fn := PyObjectFromInst(cpy3.PyCFunction_New(uintptr(unsafe.Pointer(methodCallDef)), PyNil.instance))
//	defer fn.DecRef()
//	method := PyObjectFromInst(cpy3.PyInstanceMethod(fn.instance))
//	defer method.DecRef()
//	pClass.SetAttrString(fnName, method)
//	return PyClassFromObj(pClass)
//}
