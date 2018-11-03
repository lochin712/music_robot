package util

import (
	"io/ioutil"
	"strconv"
	"syscall"
)

func SetProcessID(name string) int {
	pid := syscall.Getpid()
	if pid != 1 {
		ioutil.WriteFile(name, []byte(strconv.Itoa(pid)), 0777)
	}
	return pid
}
