package py3_go

import "runtime"

func setFinalizer(obj interface{}, finalizer interface{}) {
	runtime.SetFinalizer(obj, finalizer)
}
