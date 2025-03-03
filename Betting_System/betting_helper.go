package betting_system

import (
	"context"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomFastestLapBet(userID int64 , SessionID int) FastestLapBet {
	driverIDs := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Example driver IDs
	return FastestLapBet{
		UserID:   userID,
		SessionID:   SessionID,
		DriverID: driverIDs[rand.Intn(len(driverIDs))],
		Status:   "PENDING",
		BettingPool: 1000,
	}
}

func GenerateRandomFastestLapBets(ctx context.Context, conn *pgx.Conn, userID int64, SessionID int, count int) error {
	for i := 0; i < count; i++ {
		bet := GenerateRandomFastestLapBet(userID, SessionID)
		_, err := CreateFastestLapBet(ctx, conn, bet)
		if err != nil {
			return err
		}
	}
	return nil
}
