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
	user_management.UpdateUserTable(ctx, conn)

}
// Add a new column 'balance' with default value 0.00 to the users table

