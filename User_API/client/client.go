package client

import (
	"context"
	"f1betting/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserManagementClient struct {
	client proto.UserManagementClient
}

func NewUserManagementClient(address string) (*UserManagementClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := proto.NewUserManagementClient(conn)
	return &UserManagementClient{client: client}, nil
}

func (c *UserManagementClient) CreateUser(ctx context.Context, user *proto.User) (int64, error) {
	resp, err := c.client.CreateUser(ctx, &proto.CreateUserRequest{
		User: user,
	})
	if err != nil {
		return 0, err
	}
	return resp.UserId, nil
}

func (c *UserManagementClient) GetUserByID(ctx context.Context, id int64) (*proto.User, error) {
	return c.client.GetUserByID(ctx, &proto.GetUserByIDRequest{Id: id})
}

// Implement other client methods...

func StringPtr(s string) *wrapperspb.StringValue {
	return wrapperspb.String(s)
}

func TimePtr(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
