package graph

import (
	"f1betting/proto"
)

type Resolver struct {
	UserClient    proto.UserManagementClient
	BettingClient proto.BettingServiceClient
}
