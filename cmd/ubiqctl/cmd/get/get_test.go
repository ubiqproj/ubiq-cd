package get

import (
	"bytes"
	"context"
	"testing"
	apiv1 "ubiq-cd/third_party/connect/gen/api/v1"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/mock"
)

type MockObject struct {
	mock.Mock
}

func (m *MockObject) Greet(ctx context.Context, req *connect.Request[apiv1.GreetRequest]) (*connect.Response[apiv1.GreetResponse], error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*connect.Response[apiv1.GreetResponse]), args.Error(1)
}

func TestNewCmdGet(t *testing.T) {
	var buf bytes.Buffer

	mock := new(MockObject)
	mock.On(
		"Greet",
		context.Background(),
		connect.NewRequest(&apiv1.GreetRequest{Name: "Ubiq"}),
	).Return(
		&connect.Response[apiv1.GreetResponse]{
			Msg: &apiv1.GreetResponse{Greeting: "Hello, Ubiq!"},
		},
		nil,
	).Once()

	cmd := NewCmdGet(mock, &buf)
	if err := cmd.Execute(); err != nil {
		t.Errorf("Cannot execute version command: %v", err)
	}
	mock.AssertNumberOfCalls(t, "Greet", 1)
	mock.AssertExpectations(t)
}

func TestRunGet(t *testing.T) {
	mock := new(MockObject)
	mock.On(
		"Greet",
		context.Background(),
		connect.NewRequest(&apiv1.GreetRequest{Name: "Ubiq"}),
	).Return(
		&connect.Response[apiv1.GreetResponse]{
			Msg: &apiv1.GreetResponse{Greeting: "Hello, Ubiq!"},
		},
		nil,
	).Once()

	tests := []struct {
		name    string
		wantOut string
		wantErr bool
	}{
		{
			name:    "valid",
			wantOut: "Hello, Ubiq!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := RunGet(mock, out); (err != nil) != tt.wantErr {
				t.Errorf("RunGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("RunGet() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
	mock.AssertNumberOfCalls(t, "Greet", 1)
	mock.AssertExpectations(t)
}
