package cpy3

func PyEval_GetBuiltins() uintptr {
	r, _, _ := pyEval_GetBuiltins.Call()
	return r
}
func PyEval_GetLocals() uintptr {
	r, _, _ := pyEval_GetLocals.Call()
	return r
}

func PyEval_GetGlobals() uintptr {
	r, _, _ := pyEval_GetGlobals.Call()
	return r
}
func PyEval_GetFrame() uintptr {
	r, _, _ := pyEval_GetFrame.Call()
	return r
}
