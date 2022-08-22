package cpy3

import (
	"github.com/aadog/msvcrt-go"
	"unsafe"
)

func Py_Initialize() {
	py_Initialize.Call()
}
func Py_InitializeEx(initsigs int) {
	py_Initialize.Call(uintptr(initsigs))
}
func Py_IsInitialized() int {
	r, _, _ := py_IsInitialized.Call()
	return int(r)
}
func Py_Finalize() {
	py_Finalize.Call()
}
func Py_FinalizeEx() int {
	r, _, _ := py_FinalizeEx.Call()
	return int(r)
}

func Py_DecodeLocale(arg string, size uintptr) string {
	r, _, _ := py_DecodeLocale.Call(msvcrt.StringToCUTF8String(arg), size)
	defer PyMem_RawFree(r)
	return msvcrt.CUTF16PtrToString(r)
}
func PyMem_Free(obj uintptr) {
	pyMem_Free.Call(obj)
}
func PyMem_RawFree(obj uintptr) {
	pyMem_RawFree.Call(obj)
}

func Py_Main(args []string) int {
	argc := len(args)
	argv := make([]uintptr, 0)
	for _, arg := range args {
		argv = append(argv, msvcrt.StringToCUTF8String(arg))
	}
	r, _, _ := py_BytesMain.Call(uintptr(argc), uintptr(unsafe.Pointer(&argv[0])))
	return int(r)
}
func Py_BytesMain(args []string) int {
	argc := len(args)
	r, _, _ := py_BytesMain.Call(uintptr(argc))
	return int(r)
}

func PyRun_AnyFile(fp uintptr, filename string) int {
	r, _, _ := pyRun_AnyFile.Call(uintptr(fp), msvcrt.StringToCUTF8String(filename))
	return int(r)
}
func PyRun_SimpleFile(fp uintptr, filename string) int {
	r, _, _ := pyRun_SimpleFile.Call(uintptr(fp), msvcrt.StringToCUTF8String(filename))
	return int(r)
}
func PyRun_SimpleString(command string) int {
	r, _, _ := pyRun_SimpleString.Call(msvcrt.StringToCUTF8String(command))
	return int(r)
}

func Py_SetProgramName(name string) {
	py_SetProgramName.Call(msvcrt.StringToCUTF16String(name))
}
func Py_GetProgramName() string {
	r, _, _ := py_GetProgramName.Call()
	return msvcrt.CUTF16PtrToString(r)
}

func Py_SetPath(path string) {
	py_SetPath.Call(msvcrt.StringToCUTF16String(path))
}

func Py_SetPythonHome(home string) {
	py_SetPythonHome.Call(msvcrt.StringToCUTF16String(home))
}

func Py_fopen_obj(path uintptr, mode string) uintptr {
	r, _, _ := _py_fopen_obj.Call(path, msvcrt.StringToCUTF8String(mode))
	return r
}

func Py_IncRef(obj uintptr) {
	py_IncRef.Call(obj)
}
func Py_DecRef(obj uintptr) {
	py_DecRef.Call(obj)
}

func PyImport_AppendInittab(name string, initfunc uintptr) int {
	r, _, _ := pyImport_AppendInittab.Call(msvcrt.StringToCUTF8String(name), initfunc)
	return int(r)
}
