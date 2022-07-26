package cpy3

import (
	"syscall"
)

var (
	py_NoneStruct = libpython3.NewProc("_Py_NoneStruct")

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
	_py_fopen_obj             = libpython3.NewProc("_Py_fopen_obj")

	pyImport_AppendInittab = libpython3.NewProc("PyImport_AppendInittab")

	py_IncRef = libpython3.NewProc("Py_IncRef")
	py_DecRef = libpython3.NewProc("Py_DecRef")

	pyUnicode_AsUTF8          = libpython3.NewProc("PyUnicode_AsUTF8")
	pyUnicode_DecodeFSDefault = libpython3.NewProc("PyUnicode_DecodeFSDefault")
	pyUnicode_FromString      = libpython3.NewProc("PyUnicode_FromString")
	pyUnicode_GetLength       = libpython3.NewProc("PyUnicode_GetLength")

	pyObject_Str           = libpython3.NewProc("PyObject_Str")
	pyObject_GetAttr       = libpython3.NewProc("PyObject_GetAttr")
	pyObject_Call          = libpython3.NewProc("PyObject_Call")
	pyObject_CallObject    = libpython3.NewProc("PyObject_CallObject")
	pyObject_CallNoArgs    = libpython3.NewProc("PyObject_CallNoArgs")
	pyObject_Type          = libpython3.NewProc("PyObject_Type")
	pyObject_HasAttrString = libpython3.NewProc("PyObject_HasAttrString")
	pyObject_GetAttrString = libpython3.NewProc("PyObject_GetAttrString")
	pyObject_SetAttrString = libpython3.NewProc("PyObject_SetAttrString")
	pyObject_DelAttrString = libpython3.NewProc("PyObject_DelAttrString")

	pyLong_AsDouble     = libpython3.NewProc("PyLong_AsDouble")
	pyLong_AsLong       = libpython3.NewProc("PyLong_AsLong")
	pyLong_AsLongLong   = libpython3.NewProc("PyLong_AsLongLong")
	pyLong_FromDouble   = libpython3.NewProc("PyLong_FromDouble")
	pyLong_FromLong     = libpython3.NewProc("PyLong_FromLong")
	pyLong_FromLongLong = libpython3.NewProc("PyLong_FromLongLong")
	pyLong_FromString   = libpython3.NewProc("PyLong_FromString")

	pyModule_Create2           = libpython3.NewProc("PyModule_Create2")
	pyImport_Import            = libpython3.NewProc("PyImport_Import")
	pyModule_GetName           = libpython3.NewProc("PyModule_GetName")
	pyModule_GetDict           = libpython3.NewProc("PyModule_GetDict")
	pyModule_AddFunctions      = libpython3.NewProc("PyModule_AddFunctions")
	pyModule_AddIntConstant    = libpython3.NewProc("PyModule_AddIntConstant")
	pyModule_AddStringConstant = libpython3.NewProc("PyModule_AddStringConstant")
	pyModule_AddObject         = libpython3.NewProc("PyModule_AddObject")
	pyModule_AddObjectRef      = libpython3.NewProc("PyModule_AddObjectRef")

	pyDict_SetItem       = libpython3.NewProc("PyDict_SetItem")
	pyDict_SetItemString = libpython3.NewProc("PyDict_SetItemString")
	pyDict_New           = libpython3.NewProc("PyDict_New")
	pyDict_Size          = libpython3.NewProc("PyDict_Size")
	pyDict_Clear         = libpython3.NewProc("PyDict_Clear")
	pyDict_GetItem       = libpython3.NewProc("PyDict_GetItem")
	pyDict_Keys          = libpython3.NewProc("PyDict_Keys")
	pyDict_GetItemString = libpython3.NewProc("PyDict_GetItemString")
	pyDict_DelItem       = libpython3.NewProc("PyDict_DelItem")
	pyDict_DelItemString = libpython3.NewProc("PyDict_DelItemString")

	pyEval_GetBuiltins = libpython3.NewProc("PyEval_GetBuiltins")
	pyEval_GetLocals   = libpython3.NewProc("PyEval_GetLocals")
	pyEval_GetGlobals  = libpython3.NewProc("PyEval_GetGlobals")
	pyEval_GetFrame    = libpython3.NewProc("PyEval_GetFrame")
	pyFrame_GetBack    = libpython3.NewProc("PyFrame_GetBack")
	pyFrame_GetCode    = libpython3.NewProc("PyFrame_GetCode")

	pyTuple_Size     = libpython3.NewProc("PyTuple_Size")
	pyTuple_GetItem  = libpython3.NewProc("PyTuple_GetItem")
	pyTuple_SetItem  = libpython3.NewProc("PyTuple_SetItem")
	pyTuple_GetSlice = libpython3.NewProc("PyTuple_GetSlice")
	pyTuple_Check    = libpython3.NewProc("PyTuple_Check")
	pyTuple_New      = libpython3.NewProc("PyTuple_New")

	pyBytes_FromString = libpython3.NewProc("PyBytes_FromString")

	pyObject_CallFunctionObjArgs = libpython3.NewProc("PyObject_CallFunctionObjArgs")

	pyType_Type      = libpython3.NewProc("PyType_Type")
	pyType_Name      = libpython3.NewProc("_PyType_Name")
	pyType_GetModule = libpython3.NewProc("PyType_GetModule")

	pyList_New      = libpython3.NewProc("PyList_New")
	pyList_SetItem  = libpython3.NewProc("PyList_SetItem")
	pyList_GetItem  = libpython3.NewProc("PyList_GetItem")
	pyList_GetSlice = libpython3.NewProc("PyList_GetSlice")
	pyList_Size     = libpython3.NewProc("PyList_Size")
	pyList_Insert   = libpython3.NewProc("PyList_Insert")
	pyList_Append   = libpython3.NewProc("PyList_Append")

	pyBool_FromLong = libpython3.NewProc("PyBool_FromLong")

	pyExc_Exception  = libpython3.NewProc("PyExc_Exception")
	pyExc_ValueError = libpython3.NewProc("PyExc_ValueError")
	pyErr_SetString  = libpython3.NewProc("PyErr_SetString")

	pyInstanceMethod_New = libpython3.NewProc("PyInstanceMethod_New")

	pyCFunction_New = libpython3.NewProc("PyCFunction_New")
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
