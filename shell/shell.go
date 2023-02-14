package shell

/*
Copyright © 2023 leig <leigme@gmail.com>

*/

import (
	"fmt"
	"github.com/leigme/loki"
	"os/exec"
	"runtime"
	"strings"
)

type Shell interface {
	Exe(command string) string
	Pwd() string
}

type shell struct {
	cmdHeaders []string
	pathHeader string
	out        func(data []byte) string
}

func New() Shell {
	s := shell{}
	s.out = func(data []byte) string {
		return string(data)
	}
	if strings.EqualFold(runtime.GOOS, loki.WindowsOs) {
		s.cmdHeaders = []string{loki.WindowsCmd, "/C"}
		s.pathHeader = loki.WindowsCd
	} else {
		s.cmdHeaders = []string{loki.UnixBash, "-c"}
		s.pathHeader = loki.UnixPwd
	}
	return &s
}

func (s *shell) Exe(command string) string {
	output, err := execute(s.cmdHeaders[0], s.cmdHeaders[1], command)
	if err != nil {
		return fmt.Sprintf("execute cmd: %s, is error: %s", command, err.Error())
	}
	return s.out(output)
}

func (s *shell) Pwd() string {
	output, err := execute(s.cmdHeaders[0], s.cmdHeaders[1], s.pathHeader)
	if err != nil {
		return fmt.Sprintf("execute pwd: %s, is error: %s", s.pathHeader, err.Error())
	}
	return s.out(output)
}

func execute(args ...string) ([]byte, error) {
	cmd := exec.Command(args[0], args[1], args[2])
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
