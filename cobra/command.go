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
	cobra.Command
	Skip int
}

type Exec func(cmd *cobra.Command, args []string)

type Option func(option *CommandOption)

func NewCommand(c Command, opts ...Option) *cobra.Command {
	co := newDefaultOption()
	for _, apply := range opts {
		apply(co)
	}
	cc := &cobra.Command{
		Use:   file.Name(co.Skip),
		Short: co.Short,
		Long:  co.Long,
		Args:  co.Args,
		Run:   c.Execute(),
	}
	return cc
}

func Add(rootCmd *cobra.Command, c Command, opts ...Option) {
	rootCmd.AddCommand(NewCommand(c, opts...))
}

func AddFlags(rootCmd *cobra.Command, c Command, p *string, name, shorthand string, value string, usage string, opts ...Option) {
	cc := NewCommand(c, opts...)
	rootCmd.AddCommand(cc)
	cc.Flags().StringVarP(p, name, shorthand, value, usage)
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

func WithArgs(args cobra.PositionalArgs) Option {
	return func(option *CommandOption) {
		option.Args = args
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
	option := CommandOption{}
	option.Short = ""
	option.Long = ""
	option.Skip = 2
	return &option
}
