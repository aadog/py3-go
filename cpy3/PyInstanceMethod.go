package cpy3

func PyInstanceMethod(cfn uintptr) uintptr {
	r, _, _ := pyInstanceMethod_New.Call(cfn)
	return r
}

func PyCFunction_New(def uintptr, args uintptr) uintptr {
	r, _, _ := pyCFunction_New.Call(def, args)
	return r
}
