package cpy3

import "syscall"

var (
	py_Initialize    = libpython3.NewProc("Py_Initialize")
	py_InitializeEx  = libpython3.NewProc("Py_InitializeEx")
	py_IsInitialized = libpython3.NewProc("Py_IsInitialized")
	py_Finalize      = libpython3.NewProc("Py_Finalize")
	py_FinalizeEx    = libpython3.NewProc("Py_FinalizeEx")
	py_DecodeLocale  = libpython3.NewProc("Py_DecodeLocale")
	pyMem_Free       = libpython3.NewProc("PyMem_Free")
	pyMem_RawFree    = libpython3.NewProc("PyMem_RawFree")

	py_SetProgramName         = libpython3.NewProc("Py_SetProgramName")
	py_GetProgramName         = libpython3.NewProc("Py_GetProgramName")
	pyEval_ThreadsInitialized = libpython3.NewProc("PyEval_ThreadsInitialized")
	py_Main                   = libpython3.NewProc("Py_Main")
	py_BytesMain              = libpython3.NewProc("Py_BytesMain")
	pyRun_AnyFile             = libpython3.NewProc("PyRun_AnyFile")
	pyRun_SimpleFile          = libpython3.NewProc("PyRun_SimpleFile")
	pyRun_SimpleString        = libpython3.NewProc("PyRun_SimpleString")
	py_SetPath                = libpython3.NewProc("Py_SetPath")
	py_SetPythonHome          = libpython3.NewProc("Py_SetPythonHome")
	_py_fopen_obj          = libpython3.NewProc("_Py_fopen_obj")



	pyImport_AppendInittab          = libpython3.NewProc("PyImport_AppendInittab")
	pyModule_Create2          = libpython3.NewProc("PyModule_Create2")





	py_IncRef          = libpython3.NewProc("Py_IncRef")
	py_DecRef          = libpython3.NewProc("Py_DecRef")

	pyUnicode_FromString          = libpython3.NewProc("PyUnicode_FromString")
	pyUnicode_GetLength          = libpython3.NewProc("PyUnicode_GetLength")
	pyLong_FromLong          = libpython3.NewProc("PyLong_FromLong")
)
var kernel32dll = syscall.NewLazyDLL("kernel32.dll")
var (
	_lstrlenW = kernel32dll.NewProc("lstrlenW")
	_lstrlen  = kernel32dll.NewProc("lstrlenA")
)

var msvcrtdll = syscall.NewLazyDLL("msvcrt.dll")
var (
	_memcpy = msvcrtdll.NewProc("memcpy")
)
