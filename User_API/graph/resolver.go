package graph

import (
	"f1betting/user_api/client"
)

type Resolver struct {
	UserClient *client.UserManagementClient
}

func NewResolver(userServiceAddr string) (*Resolver, error) {
	userClient, err := client.NewUserManagementClient(userServiceAddr)
	if err != nil {
		return nil, err
	}
	return &Resolver{UserClient: userClient}, nil
}
