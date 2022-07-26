package py3

import (
	"fmt"
	"github.com/aadog/py3-go/cpy3"
	"syscall"
	"unsafe"
)

type PyClass struct {
	PyObject
}

func (p *PyClass) GetName() string {
	return p.GetAttrString("__className__").AsUTF8()
}

// PyClassFromInst
// 新建一个对象来自已经存在的对象实例指针。
//
// Create a new object from an existing object instance pointer.
func PyClassFromInst(inst uintptr) *PyClass {
	dl := new(PyClass)
	dl.instance = inst
	dl.ptr = unsafe.Pointer(dl.instance)
	return dl
}
func PyClassFromObj(obj *PyObject) *PyClass {
	dl := new(PyClass)
	dl.PyObject = *obj
	return dl
}

func CreateClass(name string, dict map[string]any) *PyClass {
	if dict == nil {
		dict = map[string]any{}
	}
	pClassName := GoToPyObject(name)
	defer pClassName.DecRef()
	pClassBases := NewPyTuple(0)
	defer pClassBases.DecRef()
	pClassDic := GoToPyObject(dict)
	defer pClassDic.DecRef()
	pClass := PyObjectFromInst(cpy3.PyObject_CallFunctionObjArgs(cpy3.PyType_Type(), pClassName.Instance(), pClassBases.Instance(), pClassDic.Instance(), 0))
	pClass.SetAttrString("__className__", NewPyUnicode(name).AsObj())

	fnName := "Call"
	methodCallDef := &cpy3.PyMethodDef{
		Ml_name: cpy3.GoStrToCStr(fnName),
		Ml_meth: syscall.NewCallback(func() uintptr {
			fmt.Println("call")

			return Py_RETURN_NONE().instance
		}),
		Ml_flags: 0,
		Ml_doc:   cpy3.GoStrToCStr("module call forward"),
	}
	PyMethodMap.Store("x", methodCallDef)

	fn := PyObjectFromInst(cpy3.PyCFunction_New(uintptr(unsafe.Pointer(methodCallDef)), cpy3.PyTuple_New(0)))
	defer fn.DecRef()
	method := PyObjectFromInst(cpy3.PyInstanceMethod(fn.instance))
	defer method.DecRef()
	pClass.SetAttrString(fnName, method)
	return PyClassFromObj(pClass)
}
