//go:build (windows || linux || darwin) && tempdll
// +build windows linux darwin
// +build tempdll

package cpy3

func checkAndReleaseDLL() (bool, string) {
	return false, ""
}
