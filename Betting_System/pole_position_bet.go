package betting_system

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type PolePositionBet struct {
	ID          int64
	UserID      int64
	SessionID      int64
	DriverID    int64
	Status      string
	BettingPool int64
	CreateAt    time.Time
}

func CreatePolePositionBet(ctx context.Context, conn *pgx.Conn, bet PolePositionBet) (int64, error) {
	var betID int64
	err := conn.QueryRow(ctx, (*bettingQueries)["CreatePolePositionBet"],
		bet.UserID, bet.SessionID, bet.DriverID, bet.BettingPool).Scan(&betID)
	return betID, err
}
