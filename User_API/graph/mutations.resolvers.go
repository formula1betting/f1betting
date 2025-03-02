package graph

import (
	"context"
	"f1betting/proto"
	"f1betting/user_api/graph/model"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func parseDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(fmt.Errorf("invalid date format: %v", err))
	}
	return date
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (string, error) {
	user := &proto.User{
		FullName:          input.FullName,
		Email:             input.Email,
		Username:          input.Username,
		PasswordHash:      input.Password,
		DateOfBirth:       timestamppb.New(parseDate(input.DateOfBirth)),
		PhoneNumber:       wrapperspb.String(*input.PhoneNumber),
		GovernmentId:      input.GovernmentID,
		Address:           input.Address,
		AccountStatus:     "ACTIVE",
		Role:              "USER",
		EmailVerified:     false,
		Country:           wrapperspb.String(*input.Country),
		PreferredCurrency: wrapperspb.String(*input.PreferredCurrency),
		FavoriteTeam:      wrapperspb.String(*input.FavoriteTeam),
		ProfilePictureUrl: wrapperspb.String(*input.ProfilePictureURL),
		Balance:           0.0,
	}

	req := &proto.CreateUserRequest{
		User: user,
	}
	resp, err := r.UserClient.CreateUser(ctx, req)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(resp.UserId, 10), nil
}

// UpdateUserProfile is the resolver for the updateUserProfile field.
func (r *mutationResolver) UpdateUserProfile(ctx context.Context, userID string, input model.UserProfileUpdateInput) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdateUserProfile - updateUserProfile"))
}

// UpdateUserEmail is the resolver for the updateUserEmail field.
func (r *mutationResolver) UpdateUserEmail(ctx context.Context, userID string, email string) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdateUserEmail - updateUserEmail"))
}

// UpdateUserPassword is the resolver for the updateUserPassword field.
func (r *mutationResolver) UpdateUserPassword(ctx context.Context, userID string, newPassword string) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdateUserPassword - updateUserPassword"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

func (r *mutationResolver) CreateFastestLapBet(ctx context.Context, userID string, input model.FastestLapBetInput) (string, error) {
	uid, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid userID: %v", err)
	}

	req := &proto.FastestLapBetRequest{
		UserId:      uid,
		SessionId:   input.SessionID,
		DriverId:    int32(input.DriverID),
		BettingPool: int32(input.BettingPool),
	}

	resp, err := r.BettingClient.CreateFastestLapBet(ctx, req)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", resp.BetId), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
