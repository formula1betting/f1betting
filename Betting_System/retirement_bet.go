package betting_system

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type RetirementBet struct {
	ID          int64
	UserID      int64
	SessionID      int64
	DriverID    int64
	Status      string
	BettingPool int64
	CreateAt    time.Time
}

func CreateRetirementBet(ctx context.Context, conn *pgx.Conn, bet RetirementBet) (int64, error) {
	var betID int64
	err := conn.QueryRow(ctx, (*bettingQueries)["CreateRetirementBet"],
		bet.UserID, bet.SessionID, bet.DriverID, bet.BettingPool).Scan(&betID)
	return betID, err
}
