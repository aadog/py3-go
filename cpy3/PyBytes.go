package cpy3

func PyBytes_FromString(s string) uintptr {
	r, _, _ := pyBytes_FromString.Call(GoStrToCStr(s))
	return r
}
