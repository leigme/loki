package cobra

import (
	"github.com/leigme/loki/file"
	"github.com/spf13/cobra"
	"strings"
)

/*
Copyright © 2023 leig <leigme@gmail.com>

*/

type Command interface {
	Execute() Exec
}

type CommandOption struct {
	Short, Long string
	Skip        int
}

type Exec func(cmd *cobra.Command, args []string)

type Option func(option *CommandOption)

func Add(rootCmd *cobra.Command, c Command, ops ...Option) {
	co := newDefaultOption()
	for _, apply := range ops {
		apply(co)
	}
	cc := &cobra.Command{
		Use:   file.Name(co.Skip),
		Short: co.Short,
		Long:  co.Long,
		Run:   c.Execute(),
	}
	rootCmd.AddCommand(cc)
}

func WithShort(short string) Option {
	return func(option *CommandOption) {
		if strings.EqualFold(short, "") {
			return
		}
		option.Short = short
	}
}

func WithLong(long string) Option {
	return func(option *CommandOption) {
		if strings.EqualFold(long, "") {
			return
		}
		option.Long = long
	}
}

func WithSkip(skip int) Option {
	return func(option *CommandOption) {
		if skip < 0 {
			return
		}
		option.Skip = skip
	}
}

func newDefaultOption() *CommandOption {
	return &CommandOption{
		Short: "",
		Long:  "",
		Skip:  2,
	}
}
