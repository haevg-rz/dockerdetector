package dockerdetector

import (
	"embed"
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
			wantErr: false,
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

func Test_isRunningInContainer(t *testing.T) {

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Running in docker",
			args: args{
				filename: "testdata/cgroup_docker",
			},
			want:    true,
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

			got, err := isRunningInContainer(f)
			if (err != nil) != tt.wantErr {
				t.Errorf("isRunningInContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isRunningInContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}
