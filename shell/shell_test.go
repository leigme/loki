package shell

import (
	"reflect"
	"testing"
)

/*
Copyright © 2023 leig HERE <leigme@gmail.com>

*/

func Test_execute(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test_execute_1",
			args: args{
				args: []string{"/bin/bash", "-c", "sleep 5s"},
			},
			want:    make([]byte, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := execute(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shell_ParseArgs(t *testing.T) {
	type fields struct {
		cmdHeaders   []string
		pathHeader   string
		out          func(data []byte) string
		shellOptions shellOptions
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult map[string]string
		wantErr    bool
	}{
		{
			name: "test_parse_args_1",
			args: args{args: []string{"-p", "8080", "-e=dev", "-d", "/users/leig"}},
			wantResult: map[string]string{
				"p": "8080",
				"e": "dev",
				"d": "/users/leig",
			},
			wantErr: false,
		},
		{
			name:    "test_parse_args_2",
			args:    args{args: []string{"-p", "8080", "-e=dev", "-d", "/users/leig", "-s"}},
			wantErr: true,
		},
		{
			name:    "test_parse_args_3",
			args:    args{args: []string{"-p", "8080", "-e=dev", "-d", "/users/leig", "-s", "-v"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shell{
				cmdHeaders:   tt.fields.cmdHeaders,
				pathHeader:   tt.fields.pathHeader,
				out:          tt.fields.out,
				shellOptions: tt.fields.shellOptions,
			}
			gotResult, err := s.ParseArgs(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ParseArgs() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
