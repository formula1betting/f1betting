package main

import (
	"context"
	"f1betting/betting_system"
	"f1betting/dbms"
	"f1betting/internal_api"
	"f1betting/user_api"
	"f1betting/user_management"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()
	conn, err := dbms.ConnectToDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbms.CloseConnection(ctx, conn)

	user_management.Init()
	betting_system.Init()

	conn.Exec(ctx, "DEALLOCATE PREPARE ALL")

	internal_api.RegisterInternalAPIs(conn)
	user_api.RegisterUserAPIs(conn)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
