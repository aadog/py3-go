package py3_go

import (
	"github.com/aadog/py3-go/cpy3"
	"reflect"
	"sync"
	"unsafe"
)

var ModuleDef = sync.Map{}

func RegPyModule(moduleName string, structV any) *PyObject {
	val := reflect.ValueOf(structV)
	inval := reflect.Indirect(val)
	tp := val.Type()
	if tp.Kind() != reflect.Pointer {
		return nil
	}
	if inval.Kind() != reflect.Struct {
		return nil
	}
	doc := ""
	docV := inval.FieldByName("Doc")
	if docV.Kind() == reflect.String {
		doc = docV.String()
	}

	methods := make([]cpy3.PyMethodDef, 0)
	for i := 0; i < tp.NumMethod(); i++ {
		method := tp.Method(i)
		methodVal := val.Method(i)
		fn, fnOk := methodVal.Interface().(PyMethodType)
		if !fnOk {
			break
		}
		methodName := method.Name
		methods = append(methods, cpy3.PyMethodDef{
			Ml_name:  cpy3.GoStrToCStr(methodName),
			Ml_meth:  NewMethodCallBack(moduleName, methodName, fn),
			Ml_flags: 1,
			Ml_doc:   cpy3.GoStrToCStr(""),
		})
	}
	methods = append(methods, cpy3.PyMethodDef{
		Ml_name:  0,
		Ml_meth:  0,
		Ml_flags: 0,
		Ml_doc:   0,
	})

	moduleDef := cpy3.PyModuleDef{
		M_base: cpy3.PyModuleDef_Base{
			Ob_base: cpy3.PyObject_HEAD_INIT(0),
		},
		M_name:     cpy3.GoStrToCStr(moduleName),
		M_doc:      cpy3.GoStrToCStr(doc),
		M_size:     -1,
		M_methods:  uintptr(unsafe.Pointer(&methods[0])),
		M_slots:    0,
		M_traverse: 0,
		M_clear:    0,
		M_free:     0,
	}
	ptr := cpy3.PyModule_Create2(uintptr(unsafe.Pointer(&moduleDef)), 1013)
	return PyObjectFromInst(ptr)
}
