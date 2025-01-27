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

	var bets []betting_system.FastestLapBet
	for i := 6; i < 10; i++ {
		bets = append(bets, betting_system.GenerateRandomFastestLapBet(int64(i), 3456))
	}

	for _, bet := range bets {
		if _, err := betting_system.CreateFastestLapBet(ctx, conn, bet); err != nil {
			log.Fatalf("Failed to create fastest lap bet: %v", err)
		}
	}

	////////////////////////////////////////////////////////////////////////////////////

	// payout, err := betting_system.GetFastestLapUserPayout(ctx, conn, 1, 1)

	// if err != nil {
	// 	log.Fatalf("Failed to get real-time odds: %v", err)
	// }

	// log.Printf("Real-time odds: %v", payout)

	// fastestLapEachLap, err := race_info.GetFastestDriverEachLap(9606)

	// if err != nil {
	// 	log.Fatalf("Failed to get fastest driver each lap: %v", err)
	// }

	// for lap, fastestLap := range fastestLapEachLap {
	// 	log.Printf("Lap %v: %v", lap, fastestLap)
	// }

	////////////////////////////////////////////////////////////////////////////////

	// fastestLap, err := race_info.GetFastestDriverWholeRace(9606)

	// if err != nil {
	// 	log.Fatalf("Failed to get fastest driver : %v", err)
	// }

	// log.Printf("Fastest lap: %v", fastestLap)

	// betting_system.GetFastestLapBetsByRace(ctx, conn, 9606, "PENDING")

	// err = betting_system.SettleBetsAndUpdateBalanceFastestLap(ctx, conn, 9606)

	// if err != nil {
	// 	log.Fatalf("Failed to settle bets: %v", err)
	// }

}
