package cpy3


func PyErr_NewException(name string,base uintptr,dict uintptr) uintptr {
	r,_,_:=pyErr_NewException.Call(GoStrToCStr(name),base,dict)
	return r
}
func PyExc_Exception() uintptr {
	return pyExc_Exception.Addr()
}

func PyExc_ValueError() uintptr {
	return pyExc_ValueError.Addr()
}

func PyErr_SetString(tp uintptr, message string) {
	pyErr_SetString.Call(tp, GoStrToCStr(message))
}
