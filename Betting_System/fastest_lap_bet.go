package betting_system

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type FastestLapBet struct {
	ID       int64
	UserID   int64
	RaceID   int64
	DriverID int64
	Amount   float64
	Status   string
	CreateAt time.Time
}

func CreateFastestLapBetTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, bettingQueries["CreateFastestLapBetsTable"])
	return err
}

func CreateFastestLapBet(ctx context.Context, conn *pgx.Conn, bet FastestLapBet) (int64, error) {
	var betID int64
	err := conn.QueryRow(ctx, bettingQueries["CreateFastestLapBet"],
		bet.UserID, bet.RaceID, bet.DriverID, bet.Amount).Scan(&betID)
	return betID, err
}

func GetFastestLapBetsByRace(ctx context.Context, conn *pgx.Conn, raceID int64, status string) ([]FastestLapBet, error) {
	rows, err := conn.Query(ctx, bettingQueries["GetRaceFastestLapBets"], raceID, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bets []FastestLapBet
	for rows.Next() {
		var bet FastestLapBet
		err := rows.Scan(&bet.ID, &bet.UserID, &bet.RaceID, &bet.DriverID, &bet.Amount, &bet.Status, &bet.CreateAt)
		if err != nil {
			return nil, err
		}
		bets = append(bets, bet)
	}
	return bets, rows.Err()
}
