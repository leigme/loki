package loki

import (
	"fmt"
	"time"
)

const (
	DirHidePrefix = "."
	WindowsOs     = "windows"
	ExeSuffix     = ".exe"
	TmpSuffix     = ".tmp"
	WindowsCd     = "cd"
	UnixPwd       = "pwd"
	WindowsCmd    = "cmd"
	UnixBash      = "/bin/bash"
)

func CostTime(t time.Time) (cost time.Duration) {
	cost = time.Since(t)
	fmt.Printf("cost time: %d\n", cost)
	return
}
