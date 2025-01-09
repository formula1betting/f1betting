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

func generateRandomFastestLapBet(userID, raceID int64) FastestLapBet {
	driverIDs := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Example driver IDs
	return FastestLapBet{
		UserID:   userID,
		RaceID:   raceID,
		DriverID: driverIDs[rand.Intn(len(driverIDs))],
		Amount:   float64(rand.Intn(1000) + 1), // Random amount between 1 and 1000
		Status:   "PENDING",
	}
}

func GenerateRandomFastestLapBets(ctx context.Context, conn *pgx.Conn, userID, raceID int64, count int) error {
	for i := 0; i < count; i++ {
		bet := generateRandomFastestLapBet(userID, raceID)
		_, err := CreateFastestLapBet(ctx, conn, bet)
		if err != nil {
			return err
		}
	}
	return nil
}
