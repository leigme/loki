package shell

/*
Copyright © 2023 leig <leigme@gmail.com>

*/

import (
	"fmt"
	"github.com/leigme/loki"
	"github.com/leigme/progressing"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Shell interface {
	Exe(command string) string
	Pwd() string
}

type shell struct {
	cmdHeaders []string
	pathHeader string
	out        func(data []byte) string
	shellOptions
}

type shellOptions struct {
	progressing.ProcessBar
}

type Option func(options *shellOptions)

func (s *shell) Exe(command string) string {
	if s.ProcessBar != nil {
		s.ProcessBar.Start()
	}
	output, err := execute(s.cmdHeaders[0], s.cmdHeaders[1], command)
	if s.ProcessBar != nil {
		s.ProcessBar.Stop()
	}
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

func New(opts ...Option) Shell {
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
	s.ProcessBar = progressing.New()
	for _, apply := range opts {
		apply(&s.shellOptions)
	}
	return &s
}

func WithProcess(bar progressing.ProcessBar) Option {
	return func(options *shellOptions) {
		if bar != nil {
			options.ProcessBar = bar
		}
	}
}

func execute(args ...string) ([]byte, error) {
	defer loki.CostTime(time.Now())
	cmd := exec.Command(args[0], args[1], args[2])
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
