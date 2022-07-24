package py3

import (
	"reflect"
)

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
	moduleDef := NewModuleDef(moduleName, doc)


	for i := 0; i < tp.NumMethod(); i++ {
		method := tp.Method(i)
		methodVal := val.Method(i)
		fn, fnOk := methodVal.Interface().(PyMethodType)
		if !fnOk {
			continue
		}
		methodName := method.Name
		moduleDef.AddMethodDef(NewMethodDef(methodName, fn, "", 1))
	}

	obj := CreateModule(moduleDef)
	return obj
}
