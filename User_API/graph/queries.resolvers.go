package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
//
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"
	"strconv"

	"f1betting/proto"
	"f1betting/user_api/graph/model"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := r.UserClient.GetUserByID(ctx, &proto.GetUserByIDRequest{Id: userID})
	if err != nil {
		return nil, err
	}

	phoneNumber := user.PhoneNumber.GetValue()
	taxId := user.TaxId.GetValue()
	country := user.Country.GetValue()
	preferredCurrency := user.PreferredCurrency.GetValue()
	favoriteTeam := user.FavoriteTeam.GetValue()
	profilePictureUrl := user.ProfilePictureUrl.GetValue()

	return &model.User{
		ID:                strconv.FormatInt(user.Id, 10),
		FullName:          user.FullName,
		Email:             user.Email,
		Username:          user.Username,
		DateOfBirth:       user.DateOfBirth.AsTime().String(),
		PhoneNumber:       &phoneNumber,
		GovernmentID:      user.GovernmentId,
		Address:           user.Address,
		TaxID:             &taxId,
		AccountStatus:     model.AccountStatus(user.AccountStatus),
		RegistrationDate:  user.RegistrationDate.AsTime().String(),
		Role:              model.UserRole(user.Role),
		EmailVerified:     user.EmailVerified,
		Country:           &country,
		PreferredCurrency: &preferredCurrency,
		FavoriteTeam:      &favoriteTeam,
		ProfilePictureURL: &profilePictureUrl,
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
	// Get all bets for session
	sessionReq := &proto.SessionRequest{
		SessionId: sessionID,
	}
	sessionBets, err := r.BettingClient.GetPendingBetsForSession(ctx, sessionReq)
	if err != nil {
		return nil, err
	}

	var result []*model.FastestLapBet
	driverPayoutMap := make(map[int32]float64)

	// Convert session bets to GraphQL model
	for _, bet := range sessionBets.Bets {
		if bet.BetType == "FASTEST_LAP" {
			result = append(result, &model.FastestLapBet{
				ID:          strconv.FormatInt(bet.Id, 10),
				UserID:      strconv.FormatInt(bet.UserId, 10),
				SessionID:   bet.SessionId,
				DriverID:    bet.GetDriver().DriverId,
				Status:      model.BetStatus(bet.Status),
				BettingPool: int32(bet.BettingPool),
				CreatedAt:   bet.CreatedAt.AsTime().String(),
			})

			// Calculate payout for each driver
			if bet.GetDriver() != nil {
				driverPayoutMap[bet.GetDriver().DriverId] += bet.Amount
			}
		}
	}

	// Convert payout map to response format
	var payouts []*model.FastestLapUserPayout
	for driverID, amount := range driverPayoutMap {
		payouts = append(payouts, &model.FastestLapUserPayout{
			DriverID: strconv.FormatInt(int64(driverID), 10),
			Payout:   amount,
		})
	}

	return &model.FastestLapBetsAndVisualizedPayout{
		FastestLapBets:   result,
		VisualizedPayout: payouts,
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
