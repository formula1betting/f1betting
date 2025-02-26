package server

import (
	"context"
	"f1betting/proto"
	"f1betting/user_management"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserManagementServer struct {
	proto.UnimplementedUserManagementServer
	conn *pgx.Conn
}

func NewUserManagementServer(conn *pgx.Conn) *UserManagementServer {
	return &UserManagementServer{conn: conn}
}

func (s *UserManagementServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	user := user_management.User{
		FullName:          req.User.FullName,
		Email:             req.User.Email,
		Username:          req.User.Username,
		PasswordHash:      req.User.PasswordHash,
		DateOfBirth:       req.User.DateOfBirth.AsTime(),
		PhoneNumber:       getStringValue(req.User.PhoneNumber),
		GovernmentID:      req.User.GovernmentId,
		Address:           req.User.Address,
		TaxID:             getStringValue(req.User.TaxId),
		Country:           getStringValue(req.User.Country),
		PreferredCurrency: getStringValue(req.User.PreferredCurrency),
		FavoriteTeam:      getStringValue(req.User.FavoriteTeam),
		ProfilePictureURL: getStringValue(req.User.ProfilePictureUrl),
		Balance:           req.User.Balance,
	}

	userID, err := user_management.CreateUser(ctx, s.conn, user)
	if err != nil {
		return nil, err
	}

	return &proto.CreateUserResponse{UserId: userID}, nil
}

func (s *UserManagementServer) GetUserByID(ctx context.Context, req *proto.GetUserByIDRequest) (*proto.User, error) {
	user, err := user_management.GetUserByID(ctx, s.conn, req.Id)
	if err != nil {
		return nil, err
	}
	return convertUserToProto(user), nil
}

// Implement other service methods similarly...

func convertUserToProto(user *user_management.User) *proto.User {
	if user == nil {
		return nil
	}

	protoUser := &proto.User{
		Id:               user.ID,
		FullName:         user.FullName,
		Email:            user.Email,
		Username:         user.Username,
		PasswordHash:     user.PasswordHash,
		DateOfBirth:      timestamppb.New(user.DateOfBirth),
		GovernmentId:     user.GovernmentID,
		Address:          user.Address,
		AccountStatus:    user.AccountStatus,
		RegistrationDate: timestamppb.New(user.RegistrationDate),
		Role:             user.Role,
		EmailVerified:    user.EmailVerified,
		Balance:          user.Balance,
	}

	if user.PhoneNumber != nil {
		protoUser.PhoneNumber = wrapperspb.String(*user.PhoneNumber)
	}
	if user.TaxID != nil {
		protoUser.TaxId = wrapperspb.String(*user.TaxID)
	}
	if user.LastPasswordChange != nil {
		protoUser.LastPasswordChange = timestamppb.New(*user.LastPasswordChange)
	}
	if user.Country != nil {
		protoUser.Country = wrapperspb.String(*user.Country)
	}
	if user.PreferredCurrency != nil {
		protoUser.PreferredCurrency = wrapperspb.String(*user.PreferredCurrency)
	}
	if user.FavoriteTeam != nil {
		protoUser.FavoriteTeam = wrapperspb.String(*user.FavoriteTeam)
	}
	if user.ProfilePictureURL != nil {
		protoUser.ProfilePictureUrl = wrapperspb.String(*user.ProfilePictureURL)
	}

	return protoUser
}

func getStringValue(wrapper *wrapperspb.StringValue) *string {
	if wrapper == nil {
		return nil
	}
	val := wrapper.GetValue()
	return &val
}


