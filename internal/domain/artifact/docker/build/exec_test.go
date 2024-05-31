package build

import (
	"reflect"
	"testing"
)

func Test_newCmdEntry(t *testing.T) {
	t.Parallel()
	type args struct {
		options []CmdEntryOptionApplier
	}
	tests := []struct {
		name      string
		args      args
		wantEntry []string
	}{
		{
			name: "with no options",
			args: args{
				options: []CmdEntryOptionApplier{},
			},
			wantEntry: []string{"docker", "build"},
		},
		{
			name: "with docker path",
			args: args{
				options: []CmdEntryOptionApplier{WithDocker("/usr/local/bin/docker")},
			},
			wantEntry: []string{"/usr/local/bin/docker", "build", "."},
		},
		{
			name: "with tag",
			args: args{
				options: []CmdEntryOptionApplier{WithTag("image:v0.1.0")},
			},
			wantEntry: []string{"docker", "build", "-t", "image:v0.1.0", "."},
		},
		{
			name: "with tag",
			args: args{
				options: []CmdEntryOptionApplier{WithTag("image1:v0.1.0"), WithTag("image2:v0.1.0")},
			},
			wantEntry: []string{"docker", "build", "-t", "image1:v0.1.0", "-t", "image2:v0.1.0", "."},
		},
		{
			name: "with build-args",
			args: args{
				options: []CmdEntryOptionApplier{WithBuildArgs(map[string]string{})},
			},
			wantEntry: []string{"docker", "build", "."},
		},
		{
			name: "with build-args",
			args: args{
				options: []CmdEntryOptionApplier{WithBuildArgs(map[string]string{
					"PORT": "8080",
				})},
			},
			wantEntry: []string{"docker", "build", "--build-args", "PORT='8080'", "."},
		},
		{
			name: "with dockerfile",
			args: args{
				options: []CmdEntryOptionApplier{withDockerfile("Dockerfile.debug")},
			},
			wantEntry: []string{"docker", "build", "-f", "Dockerfile.debug", "."},
		},
		{
			name: "with all option",
			args: args{
				options: []CmdEntryOptionApplier{
					WithDocker("/usr/local/bin/docker"),
					WithTag("image:v0.1.0"),
					WithTag("image2:v0.1.0"),
					withDockerfile("Dockerfile.debug"),
				},
			},
			wantEntry: []string{
				"/usr/local/bin/docker", "build",
				"-t", "image:v0.1.0", "-t", "image2:v0.1.0",
				"-f", "Dockerfile.debug", ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEntry := newCmdEntry(tt.args.options...); !reflect.DeepEqual(gotEntry, tt.wantEntry) {
				t.Errorf("newCmdEntry() = %v, want %v", gotEntry, tt.wantEntry)
			}
		})
	}
}
