package shell

import (
	"reflect"
	"testing"
)

/*
Copyright © 2023 leig HERE <leigme@gmail.com>

*/

func Test_shell_execute(t *testing.T) {
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
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shell{
				cmdHeaders:   tt.fields.cmdHeaders,
				pathHeader:   tt.fields.pathHeader,
				shellOptions: tt.fields.shellOptions,
			}
			got, err := s.execute(tt.args.args...)
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

func Test_shell_execute1(t *testing.T) {
	type fields struct {
		cmdHeaders   []string
		pathHeader   string
		shellOptions shellOptions
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]byte
		wantErr bool
	}{
		{
			name: "test_shell_execute_1",
			args: args{
				args: []string{"/bin/bash", "-c", "sleep 5s"},
			},
			want:    make([][]byte, 0),
			wantErr: false,
		},
		{
			name: "test_shell_execute_2",
			args: args{
				args: []string{"/bin/bash", "-c", "ls"},
			},
			want:    [][]byte{[]byte("shell.go"), []byte("shell_test.go")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shell{
				cmdHeaders:   tt.fields.cmdHeaders,
				pathHeader:   tt.fields.pathHeader,
				shellOptions: tt.fields.shellOptions,
			}
			got, err := s.execute(tt.args.args...)
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
