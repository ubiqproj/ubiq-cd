package compose

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	type args struct {
		services map[string]Service
		volumes  map[string]Volume
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				map[string]Service{
					"app": {
						Image:     "app:latest",
						Ports:     []string{"${PORT}:3000"},
						DependsOn: []string{"db"},
						Envinronment: map[string]string{
							"PORT":    "3000",
							"DB_HOST": "db",
						},
					},
					"db": {
						Image: "mysql:8.4",
						Ports: []string{"3306"},
						VolumeMounts: []VolumeMount{{
							VolumeName: "mysql",
							TargetPath: "/var/lib/mysql",
						}},
					},
				},
				map[string]Volume{
					"mysql": {},
				},
			},
			wantOut: `version: "3"
services:
  app:
    image: app:latest
    ports:
    - ${PORT}:3000
    depends_on:
    - db
    environment:
      DB_HOST: db
      PORT: "3000"
  db:
    image: mysql:8.4
    ports:
    - "3306"
    volumes:
    - mysql:/var/lib/mysql
volumes:
  mysql: {}
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.services, tt.args.volumes)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			out := &bytes.Buffer{}
			out.Write(got)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("RunVersion() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
