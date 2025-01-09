package betting_system

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type PolePositionBet struct {
	ID       int64
	UserID   int64
	RaceID   int64
	DriverID int64
	Amount   float64
	Status   string
	CreateAt time.Time
}

func CreatePolePositionBet(ctx context.Context, conn *pgx.Conn, bet PolePositionBet) (int64, error) {
	var betID int64
	err := conn.QueryRow(ctx, bettingQueries["CreatePolePositionBet"],
		bet.UserID, bet.RaceID, bet.DriverID, bet.Amount).Scan(&betID)
	return betID, err
}
