package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	dirPrefix = "."
	windowsOs = "windows"
	exeSuffix = ".exe"
)

func WorkDir() string {
	appName := Name()
	workDir := fmt.Sprint(dirPrefix, filepath.Base(appName))
	if userHome, err := os.UserHomeDir(); err == nil {
		workDir = filepath.Join(userHome, workDir)
	}
	return workDir
}

func Name() string {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	if strings.EqualFold(runtime.GOOS, windowsOs) {
		executable = strings.TrimSuffix(executable, exeSuffix)
	}
	return filepath.Base(executable)
}
