package main

import (
	"context"
	"f1betting/betting_system"
	"f1betting/dbms"
	"f1betting/user_management"
	"log"
)

func main() {
	ctx := context.Background()
	conn, err := dbms.ConnectToDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbms.CloseConnection(ctx, conn)

	user_management.Init()
	betting_system.Init()

	conn.Exec(ctx, "DEALLOCATE PREPARE ALL")

	// var bets []betting_system.FastestLapBet
	// for i := 1; i < 5; i++ {
	// 	bets = append(bets, betting_system.GenerateRandomFastestLapBet(int64(i), 1))
	// }

	// for _, bet := range bets {
	// 	if _,err := betting_system.CreateFastestLapBet(ctx, conn, bet); err != nil {
	// 		log.Fatalf("Failed to create fastest lap bet: %v", err)
	// 	}
	// }

	payout, err := betting_system.GetFastestLapUserPayout(ctx, conn, 1, 1)

	if err != nil {
		log.Fatalf("Failed to get real-time odds: %v", err)
	}

	log.Printf("Real-time odds: %v", payout)

}
