package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
//
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"
	"strconv"

	"f1betting/betting_system"
	"f1betting/user_api/graph/model"
	"f1betting/user_management"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := user_management.GetUserByID(ctx, r.Conn, userID)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:                strconv.FormatInt(user.ID, 10),
		FullName:          user.FullName,
		Email:             user.Email,
		Username:          user.Username,
		DateOfBirth:       user.DateOfBirth.String(),
		PhoneNumber:       user.PhoneNumber,
		GovernmentID:      user.GovernmentID,
		Address:           user.Address,
		TaxID:             user.TaxID,
		AccountStatus:     model.AccountStatus(user.AccountStatus),
		RegistrationDate:  user.RegistrationDate.String(),
		Role:              model.UserRole(user.Role),
		EmailVerified:     user.EmailVerified,
		Country:           user.Country,
		PreferredCurrency: user.PreferredCurrency,
		FavoriteTeam:      user.FavoriteTeam,
		ProfilePictureURL: user.ProfilePictureURL,
		Balance:           user.Balance,
	}, nil
}

// UserByEmail is the resolver for the userByEmail field.
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserByEmail - userByEmail"))
}

// UserByUsername is the resolver for the userByUsername field.
func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserByUsername - userByUsername"))
}

// FastestLapBetsAndVisualizedPayout is the resolver for the fastestLapBetsAndVisualizedPayout field.
func (r *queryResolver) FastestLapBetsAndVisualizedPayout(ctx context.Context, sessionID int32, userID string) (*model.FastestLapBetsAndVisualizedPayout, error) {
	bets, err := betting_system.GetFastestLapBetsByRace(ctx, r.Conn, int64(sessionID), "PENDING")
	if err != nil {
		return nil, err
	}

	var result []*model.FastestLapBet
	for _, bet := range *bets {
		result = append(result, &model.FastestLapBet{
			ID:          strconv.FormatInt(bet.ID, 10),
			UserID:      strconv.FormatInt(bet.UserID, 10),
			SessionID:   sessionID,
			DriverID:    int32(bet.DriverID),
			Status:      model.BetStatus(bet.Status),
			BettingPool: int32(bet.BettingPool),
			CreatedAt:   bet.CreateAt.String(),
		})
	}

	uid, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, err
	}
	visualizedPayout, err := betting_system.GetFastestLapUserVisualizedPayout(ctx, r.Conn, uid, int(sessionID))
	if err != nil {
		return nil, err
	}

	var convertedPayout []*model.FastestLapUserPayout
	for _, p := range *visualizedPayout {
		convertedPayout = append(convertedPayout, &model.FastestLapUserPayout{
			DriverID: strconv.FormatInt(int64(p.DriverID), 10),
			Payout:   p.Payout,
		})
	}

	return &model.FastestLapBetsAndVisualizedPayout{
		FastestLapBets:   result,
		VisualizedPayout: convertedPayout,
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
