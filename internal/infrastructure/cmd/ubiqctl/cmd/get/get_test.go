package get

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockObject struct {
	mock.Mock
}

func (m *MockObject) Greet(ctx context.Context, name string) (string, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(string), args.Error(1)
}

func Test_runGet(t *testing.T) {
	t.Parallel()
	m := new(MockObject)
	m.On("Greet", mock.Anything, mock.Anything).Return("Hello, Ubiq!", nil).Once()

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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := runGet(m, out); (err != nil) != tt.wantErr {
				t.Errorf("RunGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("RunGet() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
	m.AssertNumberOfCalls(t, "Greet", 1)
	m.AssertExpectations(t)
}
