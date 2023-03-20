package shell

/*
Copyright © 2023 leig <leigme@gmail.com>

*/

import (
	"bufio"
	"fmt"
	"github.com/leigme/loki"
	"github.com/leigme/loki/file"
	"github.com/leigme/progressing"
	"io"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Shell interface {
	Exe(command string) interface{}
	Pwd() string
}

type shell struct {
	cmdHeaders []string
	pathHeader string
	shellOptions
}

type shellOptions struct {
	progressing.ProcessBar
	out func(data [][]byte) interface{}
}

type Option func(options *shellOptions)

func (s *shell) Exe(command string) interface{} {
	if s.ProcessBar != nil {
		s.ProcessBar.Start()
	}
	output, err := s.execute(s.cmdHeaders[0], s.cmdHeaders[1], command)
	if s.ProcessBar != nil {
		s.ProcessBar.Stop()
	}
	if err != nil {
		return fmt.Sprintf("execute cmd: %s, is error: %s", command, err.Error())
	}
	return s.out(output)
}

func (s *shell) Pwd() string {
	output, err := s.execute(s.cmdHeaders[0], s.cmdHeaders[1], s.pathHeader)
	if err != nil {
		return fmt.Sprintf("execute pwd: %s, is error: %s", s.pathHeader, err.Error())
	}
	result := make([]string, 0)
	for _, op := range output {
		result = append(result, string(op))
	}
	return file.Arr2str(result)
}

func New(opts ...Option) Shell {
	s := shell{}
	s.out = func(data [][]byte) interface{} {
		r := make([]string, 0)
		for _, d := range data {
			r = append(r, string(d))
		}
		return r
	}
	if strings.EqualFold(runtime.GOOS, loki.WindowsOs) {
		s.cmdHeaders = []string{loki.WindowsCmd, "/C"}
		s.pathHeader = loki.WindowsCd
	} else {
		s.cmdHeaders = []string{loki.UnixBash, "-c"}
		s.pathHeader = loki.UnixPwd
	}
	for _, apply := range opts {
		apply(&s.shellOptions)
	}
	return &s
}

func WithOut(out func(data [][]byte) interface{}) Option {
	return func(options *shellOptions) {
		options.out = out
	}
}

func WithProcess(bar progressing.ProcessBar) Option {
	return func(options *shellOptions) {
		if bar != nil {
			options.ProcessBar = bar
		}
	}
}

func (s *shell) execute(args ...string) ([][]byte, error) {
	var (
		stdout    io.ReadCloser
		outputBuf *bufio.Reader
		err       error
	)
	defer func() {
		log.Printf("cmd: %s execute time is: %d\n", args[2], loki.CostTime(time.Now()))
	}()
	cmd := exec.Command(args[0], args[1], args[2])
	result := make([][]byte, 0)
	// 创建获取命令输出的管道
	if stdout, err = cmd.StdoutPipe(); err != nil {
		log.Printf("Error: Can not obtain stdout pipe: %v for cmd: %s\n", err, args[2])
		return result, err
	}
	// 执行命令
	if err = cmd.Start(); err != nil {
		log.Printf("Error: The cmd: %s start is err\n", args[2])
		return result, err
	}
	// 使用带缓冲的读取器
	outputBuf = bufio.NewReader(stdout)
	for {
		output, _, err := outputBuf.ReadLine()
		if err == io.EOF {
			break
		}
		result = append(result, output)
	}

	if err = cmd.Wait(); err != nil {
		log.Printf("Error: wait err: %v\n", err)
		return result, err
	}
	return result, nil
}
