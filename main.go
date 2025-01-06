package main

import (
	"context"
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

	// Generate and add random users
	randomUsers := user_management.GenerateRandomUsers(100)
	for _, user := range randomUsers {
		_, err := user_management.CreateUser(ctx, conn, user)
		if err != nil {
			log.Printf("Failed to create user %s: %v", user.Username, err)
			continue
		}
	}
}
