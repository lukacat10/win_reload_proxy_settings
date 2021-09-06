package main

import (
	"golang.org/x/sys/windows"
	"log"
	"syscall"
)

func main() {

	modwininet := syscall.NewLazyDLL("wininet.dll")
	iso := modwininet.NewProc("InternetSetOptionA")

	r1, _, e1 := syscall.Syscall6(
		iso.Addr(),
		4,
		0,
		39,
		0,
		0,
		0,
		0)
	if r1 == 0 {
		if e1 != windows.ERROR_SUCCESS {
			//return 0, e1
			log.Println(e1.Error())
		} else {
			//return 0, syscall.EINVAL
			log.Println(syscall.EINVAL.Error())
		}
	}
	_ = syscall.Handle(r1)
}
