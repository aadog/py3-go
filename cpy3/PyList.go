package cpy3

func PyList_New(len int64) uintptr {
	r, _, _ := pyList_New.Call(uintptr(len))
	return r
}

func PyList_SetItem(obj uintptr, index int64, item uintptr) int {
	r, _, _ := pyList_SetItem.Call(obj, uintptr(index), item)
	return int(r)
}
func PyList_GetItem(obj uintptr, index int64) uintptr {
	r, _, _ := pyList_GetItem.Call(obj, uintptr(index))
	return r
}

func PyList_GetSlice(obj uintptr, low int64, high int64) uintptr {
	r, _, _ := pyList_GetSlice.Call(obj, uintptr(low), uintptr(high))
	return r
}
func PyList_Size(obj uintptr) int64 {
	r, _, _ := pyList_Size.Call(obj)
	return int64(r)
}
func PyList_Insert(obj uintptr, index int64, item uintptr) int {
	r, _, _ := pyList_Insert.Call(obj, uintptr(index), item)
	return int(r)
}
func PyList_Append(obj uintptr, item uintptr) int {
	r, _, _ := pyList_Append.Call(obj, item)
	return int(r)
}
