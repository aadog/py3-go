package cpy3

func PyLong_AsDouble(obj uintptr) float64 {
	r, _, _ := pyLong_AsDouble.Call(obj)
	return float64(r)
}
func PyLong_AsLong(obj uintptr) int {
	r, _, _ := pyLong_AsLong.Call(obj)
	return int(r)
}
func PyLong_AsLongLong(obj uintptr) int64 {
	r, _, _ := pyLong_AsLongLong.Call(obj)
	return int64(r)
}

func PyLong_FromLong(n int) uintptr {
	r, _, _ := pyLong_FromLong.Call(uintptr(n))
	return r
}

func PyLong_FromLongLong(n int64) uintptr {
	r, _, _ := pyLong_FromLong.Call(uintptr(n))
	return r
}
func PyLong_FromDouble(n float64) uintptr {
	r, _, _ := pyLong_FromLong.Call(uintptr(n))
	return r
}

//func PyLong_FromString(n float64) uintptr {
//	r, _, _ := pyLong_FromLong.Call(uintptr(n))
//	return r
//}
