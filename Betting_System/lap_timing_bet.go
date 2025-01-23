package betting_system

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type LapTimingBet struct {
	ID          int64
	UserID      int64
	SessionID   int64
	LapNumber   int64
	DriverID    int64
	Status      string
	BettingPool int64
	CreateAt    time.Time
}

func CreateLapTimingBet(ctx context.Context, conn *pgx.Conn, bet LapTimingBet) (int64, error) {
	var betID int64
	err := conn.QueryRow(ctx, (*bettingQueries)["CreateLapTimingBet"],
		bet.UserID, bet.SessionID, bet.LapNumber, bet.DriverID, bet.BettingPool).Scan(&betID)
	return betID, err
}
