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

	bets, err := betting_system.GetFastestLapBetsByRace(ctx, conn, 1, "PENDING")
	if err != nil {
		log.Fatalf("Failed to get fastest lap bets: %v", err)
	}

	for _, bet := range bets {
		log.Printf("Bet ID: %d, User ID: %d, ", bet.ID, bet.UserID)
	}

}
