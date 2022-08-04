package cpy3

import (
	"github.com/aadog/dylib"
	"runtime"
)

var libpython3 = loadUILib()

var (
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
)

// 加载库
func loadUILib() *dylib.LazyDLL {
	libName := getDLLName()
	// 如果支持运行时释放，则使用此种方法
	if support, newDLLPath := checkAndReleaseDLL(); support {
		libName = newDLLPath
	} else {
		libName = libName
	}
	lib := dylib.NewLazyDLL(libName)
	err := lib.Load()
	if err != nil {
		panic(err)
	}

	return lib
}

func getDLLName() string {
	libName := "python38"
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}
