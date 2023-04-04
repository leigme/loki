package config

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("env.id", "123")
	os.Setenv("env.name", "lisi")
	os.Setenv("env.sex", "false")
	m.Run()
}

func TestEnv(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_Env_1",
			args: args{&EnvTestArg{
				Id:   99,
				Name: "zhansan",
				Sex:  true,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Env(tt.args.value)
			fmt.Println(tt.args.value)
		})
	}
}

type EnvTestArg struct {
	Id   int    `env:"env.id"`
	Name string `env:"env.name"`
	Sex  bool   `env:"env.sex"`
}
