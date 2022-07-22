//go:build !tempdll && !memorydll
// +build !tempdll,!memorydll

package cpy3

func checkAndReleaseDLL() (bool, string) {
	return false, ""
}
