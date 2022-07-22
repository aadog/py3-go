package cpy3

import (
	"syscall"
	"unsafe"
)

type Pointer = uintptr
type SIZE_T = uintptr

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
}

//func CStrToGoStr(ustr uintptr)string{
//	return copyStr(ustr,G_strlen(ustr))
//}

func CBytesToGoBytes(ustr uintptr, n int) []byte {
	return copyBytes(ustr, n)
}

// 这种跟copyStr3基本一样，只是用go来处理了
func copyBytes(src uintptr, strLen int) []byte {
	if strLen == 0 {
		return nil
	}
	str := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		str[i] = *(*byte)(unsafe.Pointer(src + uintptr(i)))
	}
	return str
}

// 这种跟copyStr3基本一样，只是用go来处理了
func copyStr(src uintptr, strLen int) string {
	if strLen == 0 {
		return ""
	}
	str := make([]uint8, strLen)
	for i := 0; i < strLen; i++ {
		str[i] = *(*uint8)(unsafe.Pointer(src + uintptr(i)))
	}
	return string(str)
}

// Go的string转换为Lazarus的string
func GoStrToCStr(s string) uintptr {
	if s == "" {
		return 0
	}
	return uintptr(unsafe.Pointer(StringToUTF8Ptr(s)))
}

// Go的string转换为Lazarus的string
func GoByteToCPtr(b []byte) uintptr {
	if len(b) == 0 {
		return 0
	}
	return uintptr(unsafe.Pointer(&b[0]))
}

// 字符串到UTF8指针
func StringToUTF8Ptr(s string) *uint8 {
	temp := []byte(s)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结尾
	copy(utf8StrArr, temp)
	return &utf8StrArr[0]
}

// 字符串到UTF16指针
func StringToUTF16Ptr(s string) uintptr {
	p, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(p))
}

// UTF16指针到字符串
func UTF16PtrToString(p uintptr) string {
	l := LstrlenW(p)
	if l == 0 {
		return ""
	}
	buff := make([]uint16, l)
	Memcpy(uintptr(unsafe.Pointer(&buff[0])), p, uintptr(l*2))
	return syscall.UTF16ToString(buff)
}

func LstrlenW(lpString uintptr) int32 {
	r, _, _ := _lstrlenW.Call(lpString)
	return int32(r)
}
func Lstrlen(lpString uintptr) int32 {
	r, _, _ := _lstrlen.Call(lpString)
	return int32(r)
}
func Memcpy(dest, src Pointer, count SIZE_T) Pointer {
	r, _, _ := _memcpy.Call(dest, src, count)
	return r
}
