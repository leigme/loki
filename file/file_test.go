package file

import (
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		filename string
		data     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestCreate_1",
			args: args{
				filename: "test.txt",
				data:     []byte("abcdefg"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.filename, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Delete(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestDelete_1",
			args:    args{filename: "test_delete.txt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create("test_delete.txt", []byte("test_delete.txt")); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Delete(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExist(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestExist_1",
			args: args{filename: "test_exist.txt"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create("test_exist.txt", []byte("test_exist")); err != nil {
				log.Fatal(err)
			}
			if got := Exist(tt.args.filename); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
			if err := Delete("test_exist.txt"); err != nil {
				log.Fatal(err)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	type args struct {
		filename  string
		chunkSize int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSplit_1",
			args: args{
				filename:  "test_split.txt",
				chunkSize: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create("test_split.txt", []byte("test_split")); (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Split(tt.args.filename, tt.args.chunkSize); (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Delete("test_split.txt"); err != nil {
				log.Fatal(err)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		dir      string
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestMerge_1",
			args: args{
				dir:      "./",
				filename: "test_merge.txt",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create("test_merge.txt", []byte("test_merge")); (err != nil) != tt.wantErr {
				t.Errorf("Merge() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Split("test_merge.txt", 3); (err != nil) != tt.wantErr {
				t.Errorf("Merge() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Delete("test_merge.txt"); err != nil {
				t.Errorf("Merge() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Merge(tt.args.dir, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Merge() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Delete("test_merge.txt"); err != nil {
				t.Errorf("Merge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateDir(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestCreateDir_1",
			args:    args{filename: "./abc"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDir(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("CreateDir() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Delete("./abc"); (err != nil) != tt.wantErr {
				t.Errorf("CreateDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMd5(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "TestMd5_1",
			args:    args{filename: "TestMd5.txt"},
			want:    "9050bddcf415f2d0518804e551c1be98",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Create("TestMd5.txt", []byte("test_md5"))
			if (err != nil) != tt.wantErr {
				t.Errorf("Md5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := Md5(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Md5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Md5() got = %v, want %v", got, tt.want)
			}
			err = Delete("TestMd5.txt")
			if (err != nil) != tt.wantErr {
				t.Errorf("Md5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
