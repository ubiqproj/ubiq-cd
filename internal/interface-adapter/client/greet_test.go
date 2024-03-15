package client

import (
	"context"
	"testing"
	apiv1 "ubiq-cd/internal/interface-adapter/interface/connectrpc/gen/api/v1"
	"ubiq-cd/internal/interface-adapter/interface/connectrpc/gen/api/v1/apiv1connect"

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

func Test_client_Greet(t *testing.T) {
	t.Parallel()
	m := new(MockObject)
	m.On("Greet", mock.Anything, mock.Anything).Return(
		&connect.Response[apiv1.GreetResponse]{
			Msg: &apiv1.GreetResponse{Greeting: "Hello, Ubiq!"},
		},
		nil,
	).Once()

	type fields struct {
		GreetServiceClient apiv1connect.GreetServiceClient
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				GreetServiceClient: m,
			},
			args: args{
				ctx:  context.Background(),
				name: "UbiqCD",
			},
			want:    "Hello, Ubiq!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				GreetServiceClient: tt.fields.GreetServiceClient,
			}
			got, err := c.Greet(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Greet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.Greet() = %v, want %v", got, tt.want)
			}
		})
	}
	m.AssertNumberOfCalls(t, "Greet", 1)
	m.AssertExpectations(t)
}
