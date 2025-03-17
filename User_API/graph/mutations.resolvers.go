package graph

import (
	"context"
	"f1betting/betting_system"
	"f1betting/user_api/graph/model"
	"f1betting/user_management"
	"fmt"
	"strconv"
	"time"
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
	if err := user_management.ValidateEmail(input.Email); err != nil {
		return "", err
	}

	// Hash password
	hashedPassword, err := user_management.HashPassword(input.Password)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}

	//create user object
	user := &user_management.User{
		FullName:          input.FullName,
		Email:             input.Email,
		Username:          input.Username,
		PasswordHash:      hashedPassword, // Use hashed password instead of plain text
		DateOfBirth:       parseDate(input.DateOfBirth),
		PhoneNumber:       input.PhoneNumber,
		GovernmentID:      input.GovernmentID,
		Address:           input.Address,
		TaxID:             input.TaxID,
		Country:           input.Country,
		PreferredCurrency: input.PreferredCurrency,
		FavoriteTeam:      input.FavoriteTeam,
		ProfilePictureURL: input.ProfilePictureURL,
	}

	resp, err := user_management.CreateUser(ctx, r.Conn, *user)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(resp, 10), nil
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

	bet := &betting_system.FastestLapBet{
		UserID:      uid,
		SessionID:   int(input.SessionID),
		DriverID:    input.DriverID,
		BettingPool: int64(input.BettingPool),
	}

	resp, err := betting_system.CreateFastestLapBet(ctx, r.Conn, *bet)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(resp, 10), nil

}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
