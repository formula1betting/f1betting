package betting_system

import (
	"context"
	"math"
	"time"

	"github.com/jackc/pgx/v5"
)

// Global variable to store fastest lap bets
var fastestLapBets = make([]FastestLapBet, 0)

type FastestLapBet struct {
	ID          int64
	UserID      int64
	SessionID   int
	DriverID    int32
	Amount      float64
	Status      string
	BettingPool int64
	CreateAt    time.Time
}

func (bet *FastestLapBet) SetFastestLapBet(userID int64, sessionID int, driverID int32, bettingPool int64) {
	bet.UserID = userID
	bet.SessionID = sessionID
	bet.DriverID = driverID
	bet.BettingPool = bettingPool
}

func CreateFastestLapBetTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, (*bettingQueries)["CreateFastestLapBetsTable"])
	return err
}

func CreateFastestLapBet(ctx context.Context, conn *pgx.Conn, bet FastestLapBet) (int64, error) {
	var betID int64

	err := conn.QueryRow(ctx, (*bettingQueries)["CreateFastestLapBet"],
		bet.UserID, bet.SessionID, bet.DriverID, 0.95*float64(bet.BettingPool), bet.BettingPool).Scan(&betID)

	if err == nil {
		bet.ID = betID
		bet.Status = "PENDING"
		fastestLapBets = append(fastestLapBets, bet)
	}

	return betID, err
}

func GetFastestLapBetsByRace(ctx context.Context, conn *pgx.Conn, SessionID int64, status string) (*[]FastestLapBet, error) {
	rows, err := conn.Query(ctx, (*bettingQueries)["GetSessionFastestLapBets"], SessionID, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bets []FastestLapBet
	for rows.Next() {
		var bet FastestLapBet
		err := rows.Scan(&bet.ID, &bet.UserID, &bet.SessionID, &bet.DriverID, &bet.Amount, &bet.Status, &bet.BettingPool, &bet.CreateAt)
		if err != nil {
			return nil, err
		}
		bets = append(bets, bet)
	}

	fastestLapBets = bets

	return &bets, rows.Err()
}

type FastestLapUserPayout struct {
	DriverID int32
	Payout   float64
}

// GetFastestLapUserVisualizedPayout retrieves the potential payout for a user if they win a specific bet.
func GetFastestLapUserVisualizedPayout(ctx context.Context, conn *pgx.Conn, userID int64, SessionID int) (*[]FastestLapUserPayout, error) {
	var totalPool float64
	var userBets []FastestLapBet
	driverBets := make(map[int32][]float64)

	// Use stored bets instead of querying DB
	for _, bet := range fastestLapBets {
		if bet.SessionID == SessionID && bet.Status == "PENDING" {
			totalPool += bet.Amount
			if bet.UserID == userID {
				userBets = append(userBets, bet)
			}
			driverBets[bet.DriverID] = append(driverBets[bet.DriverID], bet.Amount)
		}
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

	return &userPayouts, nil
}
