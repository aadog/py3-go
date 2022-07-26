package cpy3

func PyTuple_Size(obj uintptr) int64 {
	r, _, _ := pyTuple_Size.Call(obj)
	return int64(r)
}
func PyTuple_Check(obj uintptr) int64 {
	r, _, _ := pyTuple_Check.Call(obj)
	return int64(r)
}

func PyTuple_GetItem(obj uintptr, pos int64) uintptr {
	r, _, _ := pyTuple_GetItem.Call(obj, uintptr(pos))
	return r
}
func PyTuple_SetItem(obj uintptr, pos int64, o uintptr) uintptr {
	r, _, _ := pyTuple_SetItem.Call(obj, uintptr(pos), o)
	return r
}

func PyTuple_GetSlice(obj uintptr, low int64, high int64) uintptr {
	r, _, _ := pyTuple_GetSlice.Call(obj, uintptr(low), uintptr(high))
	return r
}

func PyTuple_New(len int64) uintptr {
	r, _, _ := pyTuple_New.Call(uintptr(len))
	return r
}
