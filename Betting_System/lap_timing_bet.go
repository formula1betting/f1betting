package betting_system

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type LapTimingBet struct {
	ID        int64
	UserID    int64
	RaceID    int64
	LapNumber int64
	DriverID  int64
	Amount    float64
	Status    string
	CreateAt  time.Time
}

func CreateLapTimingBet(ctx context.Context, conn *pgx.Conn, bet LapTimingBet) (int64, error) {
	var betID int64
	err := conn.QueryRow(ctx, bettingQueries["CreateLapTimingBet"],
		bet.UserID, bet.RaceID, bet.LapNumber, bet.DriverID, bet.Amount).Scan(&betID)
	return betID, err
}

