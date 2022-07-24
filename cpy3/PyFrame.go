package cpy3

func PyFrame_GetBack(frame uintptr) uintptr {
	r, _, _ := pyFrame_GetBack.Call(frame)
	return r
}
func PyFrame_GetCode(frame uintptr) uintptr {
	r, _, _ := pyFrame_GetCode.Call(frame)
	return r
}
