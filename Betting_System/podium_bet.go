package betting_system

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type PodiumBet struct {
	ID             int64
	UserID         int64
	SessionID         int64
	FirstPosition  string
	SecondPosition string
	ThirdPosition  string
	Status         string
	BettingPool    int64
	CreateAt       time.Time
}

func CreatePodiumBet(ctx context.Context, conn *pgx.Conn, bet PodiumBet) (int64, error) {
	var betID int64
	err := conn.QueryRow(ctx, (*bettingQueries)["CreatePodiumBet"],
		bet.UserID, bet.SessionID, bet.FirstPosition, bet.SecondPosition, bet.ThirdPosition, bet.BettingPool).Scan(&betID)
	return betID, err
}