package betting_system

import (
	"context"
	"math"
	"time"

	"github.com/jackc/pgx/v5"
)

type FastestLapBet struct {
	ID          int64
	UserID      int64
	RaceID      int64
	DriverID    int64
	Amount      float64
	Status      string
	BettingPool int64
	CreateAt    time.Time
}

func CreateFastestLapBetTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, bettingQueries["CreateFastestLapBetsTable"])
	return err
}

func CreateFastestLapBet(ctx context.Context, conn *pgx.Conn, bet FastestLapBet) (int64, error) {
	var betID int64

	err := conn.QueryRow(ctx, bettingQueries["CreateFastestLapBet"],
		bet.UserID, bet.RaceID, bet.DriverID, bet.Amount, bet.BettingPool).Scan(&betID)
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
		err := rows.Scan(&bet.ID, &bet.UserID, &bet.RaceID, &bet.DriverID, &bet.Amount, &bet.Status, &bet.BettingPool, &bet.CreateAt)
		if err != nil {
			return nil, err
		}
		bets = append(bets, bet)
	}
	return bets, rows.Err()
}

type FastestLapUserPayout struct {
	DriverID int64
	Payout   float64
}

// GetUserPayout retrieves the potential payout for a user if they win a specific bet.
func GetFastestLapUserPayout(ctx context.Context, conn *pgx.Conn, userID int64, raceID int) ([]FastestLapUserPayout, error) {
	query := bettingQueries["GetRaceFastestLapBets"]

	rows, err := conn.Query(ctx, query, raceID, "PENDING")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var totalPool float64
	var userBets []FastestLapBet
	var allBets []FastestLapBet

	driverBets := make(map[int64][]float64)

	for rows.Next() {
		var bet FastestLapBet
		err := rows.Scan(&bet.ID, &bet.UserID, &bet.RaceID, &bet.DriverID, &bet.Amount, &bet.Status, &bet.BettingPool, &bet.CreateAt)
		if err != nil {
			return nil, err
		}
		totalPool += bet.Amount
		allBets = append(allBets, bet)
		if bet.UserID == userID {
			userBets = append(userBets, bet)
		}
		driverBets[bet.DriverID] = append(driverBets[bet.DriverID], bet.Amount)
	}

	if len(userBets) == 0 {
		return nil, nil
	}

	var userPayouts []FastestLapUserPayout
	for _, bet := range userBets {

		sum := float64(0)
		for _, num := range driverBets[bet.DriverID] {
			sum += num
		}

		negation := totalPool - sum

		membersInWinningPool := len(driverBets[bet.DriverID])

		userPayout := math.Floor(float64(negation)/float64(membersInWinningPool)) + bet.Amount
		userPayouts = append(userPayouts, FastestLapUserPayout{DriverID: bet.DriverID, Payout: userPayout})

	}

	return userPayouts, nil

}
