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
