package dockerdetector

import (
	"embed"
	"runtime"
	"testing"
)

//go:embed testdata/*
var testdataFS embed.FS

func TestIsRunningInContainer(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "No panic",
			wantErr: runtime.GOOS != "linux",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := IsRunningInContainer()
			if (err != nil) != tt.wantErr {
				t.Errorf("IsRunningInContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestInternalIsRunningInContainer(t *testing.T) {

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantID  string
		wantErr bool
	}{
		{
			name: "Running in docker",
			args: args{
				filename: "testdata/cgroup_docker",
			},
			want:    true,
			wantID:  "4531c6cdf6e13484be06e3615ebf4721c51a0b814555b15c210115762fc5b484",
			wantErr: false,
		},
		{
			name: "Empty cgroup",
			args: args{
				filename: "testdata/cgroup_empty",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Do not run in docker",
			args: args{
				filename: "testdata/cgroup_vps",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := testdataFS.Open(tt.args.filename)
			if err != nil {
				t.Errorf("isRunningInContainer() error = %v", err)
				return
			}

			got, id, err := isRunningInContainer(f)
			if (err != nil) != tt.wantErr {
				t.Errorf("isRunningInContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isRunningInContainer() = %v, want %v", got, tt.want)
			}
			if id != tt.wantID {
				t.Errorf(`isRunningInContainer() = '%v', want '%v'`, id, tt.wantID)
			}
		})
	}
}

func Test_protect(t *testing.T) {
	type args struct {
		appID string
		key   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No panic",
			args: args{
				appID: "This is a salt",
				key:   "572ab71a5256209699951225618140e258b2c9129a96b89573ced3480a2d8bd7",
			},
			want: "746a7896f415ac60fbe1955f76546f10d9529e24d261c0fe83dc751aef81bd06",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protect(tt.args.appID, tt.args.key); got != tt.want {
				t.Errorf("protect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createID(t *testing.T) {
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
			name: "Docker",
			args: args{
				filename: "testdata/cgroup_docker",
			},
			want:    "572ab71a5256209699951225618140e258b2c9129a96b89573ced3480a2d8bd7",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := testdataFS.Open(tt.args.filename)
			got, err := createID(f)
			if (err != nil) != tt.wantErr {
				t.Errorf("createID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createProtectedID(t *testing.T) {
	type args struct {
		salt     string
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Docker",
			args: args{
				salt:     "This is a salt",
				filename: "testdata/cgroup_docker",
			},
			want:    "746a7896f415ac60fbe1955f76546f10d9529e24d261c0fe83dc751aef81bd06",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := testdataFS.Open(tt.args.filename)
			got, err := createProtectedID(tt.args.salt, f)
			if (err != nil) != tt.wantErr {
				t.Errorf("createProtectedID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createProtectedID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateID(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "No panic",
			wantErr: runtime.GOOS != "linux",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateID()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCreateProtectedID(t *testing.T) {
	type args struct {
		salt string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "No panic",
			wantErr: runtime.GOOS != "linux",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateProtectedID(tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProtectedID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
