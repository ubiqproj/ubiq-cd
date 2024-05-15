package dockerfile

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	type args struct {
		stages []Stage
	}
	builderStage := "builder"
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "valid",
			args: args{stages: []Stage{{
				As:            &builderStage,
				Image:         "golang",
				Envinronments: map[string]string{"ROOT": "/go/src/app"},
				Copies: []Copy{{
					Src: []string{"."},
					To:  ".",
				}},
				Runs:       []string{"CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary"},
				Cmds:       []string{},
				EntryPoint: []string{},
			}, {
				Image:         "busybox",
				Envinronments: map[string]string{"ROOT": "/go/src/app"},
				Copies: []Copy{{
					From: &builderStage,
					Src:  []string{"${ROOT}/binary"},
					To:   "${ROOT}",
				}},
				Runs:       []string{},
				Cmds:       []string{},
				EntryPoint: []string{"${ROOT}/binary"},
			}}},
			wantOut: `FROM golang AS builder
ENV ROOT=/go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary

FROM busybox
ENV ROOT=/go/src/app
COPY --from=builder ${ROOT}/binary ${ROOT}
ENTRYPOINT [ "${ROOT}/binary" ]`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.stages)
			out := &bytes.Buffer{}
			out.Write(got)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("RunVersion() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
