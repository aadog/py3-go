package py3

import (
	"github.com/aadog/py3-go/cpy3"
	"sync"
)

var SystemModuleMap = sync.Map{}

func Initialize() {
	cpy3.Py_Initialize()
}
func AddModuleToSystemMap(m *PyModule) {
	SystemModuleMap.Store(m.GetName(), m)
}
func IsInitialized() int {
	return cpy3.Py_IsInitialized()
}
func Finalize() {
	SystemModuleMap.Range(func(key, value any) bool {
		SystemModuleMap.Delete(key)
		m := value.(*PyModule)
		m.DecRef()
		return true
	})
	cpy3.Py_Finalize()
}
func FinalizeEx() int {
	SystemModuleMap.Range(func(key, value any) bool {
		SystemModuleMap.Delete(key)
		m := value.(*PyModule)
		m.DecRef()
		return true
	})
	return cpy3.Py_FinalizeEx()
}
