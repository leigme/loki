package app

import (
	"fmt"
	"github.com/leigme/loki"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// WorkDir 应用运行的工作目录
func WorkDir() string {
	appName := Name()
	workDir := fmt.Sprint(loki.DirHidePrefix, filepath.Base(appName))
	if userHome, err := os.UserHomeDir(); err == nil {
		workDir = filepath.Join(userHome, workDir)
	}
	return workDir
}

// Name 应用启动的名称
func Name() string {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	if strings.EqualFold(runtime.GOOS, loki.WindowsOs) {
		executable = strings.TrimSuffix(executable, loki.ExeSuffix)
	}
	return filepath.Base(executable)
}
