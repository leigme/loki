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
	flags []Flag
	Skip  int
}

type Flag struct {
	P                             *string
	Name, Shorthand, Value, Usage string
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
	for _, flag := range co.flags {
		cc.Flags().StringVarP(flag.P, flag.Name, flag.Shorthand, flag.Value, flag.Usage)
	}
	return cc
}

func Add(rootCmd *cobra.Command, c Command, opts ...Option) {
	rootCmd.AddCommand(NewCommand(c, opts...))
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

func WithFlags(flags []Flag) Option {
	return func(option *CommandOption) {
		if len(flags) > 0 {
			option.flags = flags
		}
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
	option.Skip = 3
	return &option
}
