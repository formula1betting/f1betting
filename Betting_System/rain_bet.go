package betting_system

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type RainBet struct {
	ID             int64
	UserID         int64
	RaceID         int64
	RainPrediction bool
	Amount         float64
	Status         string
	CreateAt       time.Time
}

func CreateRainBet(ctx context.Context, conn *pgx.Conn, bet RainBet) (int64, error) {
	var betID int64
	err := conn.QueryRow(ctx, bettingQueries["CreateRainBet"],
		bet.UserID, bet.RaceID, bet.RainPrediction, bet.Amount).Scan(&betID)
	return betID, err
}
