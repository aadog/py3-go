package py3

import (
	"reflect"
)

func GoToPyObject(o interface{}) *PyObject {
	to := reflect.TypeOf(o)
	valOf := reflect.ValueOf(o)

	if o == nil {
		return Py_RETURN_NONE().AsObj()
	}
	//*PyObject return
	if to.AssignableTo(reflect.TypeOf(&PyObject{})) {
		return o.(*PyObject)
	}
	if valOf.CanInt() {
		return PyLongFromLongLong(valOf.Int()).AsObj()
	}
	if valOf.CanUint() {
		return PyLongFromLongLong(int64(valOf.Uint())).AsObj()
	}
	if valOf.CanFloat() {
		return PyLong_FromDouble(valOf.Float()).AsObj()
	}

	if to.Kind() == reflect.Bool {
		return NewPyBool(valOf.Bool())
	}
	if to.Kind() == reflect.String {
		return NewPyUnicode(valOf.String()).AsObj()
	}

	if to.Kind() == reflect.Slice {
		l := NewPyList(0)
		size := valOf.Len()
		for i := 0; i < size; i++ {
			item := valOf.Index(i).Interface()
			o := GoToPyObject(item)
			defer o.DecRef()
			l.Append(o)
		}
		return l.AsObj()
	}
	if to.Kind() == reflect.Map {
		d := NewPyDict()
		keys := valOf.MapKeys()
		for _, key := range keys {
			k := GoToPyObject(key.Interface())
			defer k.DecRef()
			v := GoToPyObject(valOf.MapIndex(key).Interface())
			defer v.DecRef()
			d.SetItem(k, v)
		}
		return d.AsObj()
	}
	return Py_RETURN_NONE().AsObj()
}

func PyObjectToGo(o *PyObject, to reflect.Type) any {
	//*PyObject return
	if to.AssignableTo(reflect.TypeOf(o)) {
		return o
	}
	if to.Kind() == reflect.Int {
		return int(o.AsLongLong())
	}
	if to.Kind() == reflect.Int8 {
		return int8(o.AsLongLong())
	}
	if to.Kind() == reflect.Int16 {
		return int16(o.AsLongLong())
	}
	if to.Kind() == reflect.Int32 {
		return int32(o.AsLongLong())
	}
	if to.Kind() == reflect.Int64 {
		return int64(o.AsLongLong())
	}

	if to.Kind() == reflect.Uint {
		return uint(o.AsLongLong())
	}
	if to.Kind() == reflect.Uint8 {
		return uint8(o.AsLongLong())
	}
	if to.Kind() == reflect.Uint16 {
		return uint16(o.AsLongLong())
	}
	if to.Kind() == reflect.Uint32 {
		return uint32(o.AsLongLong())
	}
	if to.Kind() == reflect.Uint64 {
		return uint64(o.AsLongLong())
	}

	if to.Kind() == reflect.Uintptr {
		return uintptr(o.AsLongLong())
	}

	if to.Kind() == reflect.Bool {
		return o.AsInt() != 0
	}
	if to.Kind() == reflect.String {
		return o.AsUTF8()
	}
	if to.Kind() == reflect.Float32 {
		return float32(o.AsDouble())
	}
	if to.Kind() == reflect.Float64 {
		return float64(o.AsDouble())
	}
	if to.Kind() == reflect.Slice {
		l := reflect.New(to).Elem()
		pyL := PyListFromObj(o.AsObj())
		for i := int64(0); i < pyL.Size(); i++ {
			l = reflect.Append(l, reflect.ValueOf(PyObjectToGo(pyL.GetItem(i), to.Elem())))
		}
		return l.Interface()
	}
	if to.Kind() == reflect.Map {
		l := reflect.MakeMap(to)
		pyD := PyDictFromObj(o.AsObj())
		keys := PyListFromObj(pyD.Keys())
		defer keys.DecRef()
		for i := int64(0); i < keys.Size(); i++ {
			k := keys.GetItem(i)
			goK := PyObjectToGo(k, to.Key())
			l.SetMapIndex(reflect.ValueOf(goK), reflect.ValueOf(PyObjectToGo(pyD.GetItem(k), to.Elem())))
		}
		return l.Interface()
	}
	if to.Kind() == reflect.Func {
		f := reflect.MakeFunc(to, func(args []reflect.Value) (results []reflect.Value) {
			pyArgs := NewPyTuple(int64(len(args)))
			for i, arg := range args {
				pyArgs.SetItem(int64(i), GoToPyObject(arg.Interface()))
			}
			rets := o.Call(pyArgs.AsObj(), PyNil)
			//如果声明函数不需要返回
			if to.NumOut() == 0 {
				return nil
			}
			tp := rets.Type()
			defer tp.DecRef()

			goRets := make([]reflect.Value, 0)

			//如果是单个返回
			if tp.Name() != "tuple" {
				if to.NumOut() != 1 {
					k := make([]reflect.Value, 0)
					for i := 0; i < to.NumOut(); i++ {
						k = append(k, reflect.ValueOf(reflect.New(to.Out(i)).Elem().Interface()))
					}
					PyErr_SetString(PyExc_ValueError(), "返回参数不正确")
					return k
				}
				goRets = append(goRets, reflect.ValueOf(PyObjectToGo(rets, to.Out(0))))
			} else {
				tpRets := PyTupleFromObj(rets.AsObj())
				if int64(to.NumOut()) != tpRets.Size() {
					k := make([]reflect.Value, 0)
					for i := 0; i < to.NumOut(); i++ {
						k = append(k, reflect.ValueOf(reflect.New(to.Out(i)).Elem().Interface()))
					}
					PyErr_SetString(PyExc_Exception(), "返回参数不正确")
					return k
				}
				for i := int64(0); i < tpRets.Size(); i++ {
					goRets = append(goRets, reflect.ValueOf(PyObjectToGo(tpRets.GetItem(i), to.Out(0))))
				}
			}
			return goRets
		})
		return f.Interface()
	}

	return o
}
