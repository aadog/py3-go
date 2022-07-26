package cpy3

func PyExc_Exception() uintptr {
	return pyExc_Exception.Addr()
}

func PyExc_ValueError() uintptr {
	return pyExc_ValueError.Addr()
}

func PyErr_SetString(tp uintptr, message string) {
	pyErr_SetString.Call(tp, GoStrToCStr(message))
}
