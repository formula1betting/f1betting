package dbms

import (
	"log"
	"sync"
)

var (
	onceUserManagement sync.Once
	onceBettingSystem  sync.Once
)

func GetUserManagementQueries() *map[string]string {
	var queries *map[string]string

	onceUserManagement.Do(func() {
		var err error
		queries, err = loadSQLFile("./Database/sql/user_management.sql")
		if err != nil {
			log.Fatalf("Failed to load user management queries: %v", err)
		}
	})
	return queries
}

func GetBettingQueries() *map[string]string {
	var queries *map[string]string

	onceBettingSystem.Do(func() {
		var err error
		queries, err = loadSQLFile("./Database/sql/betting_system.sql")
		if err != nil {
			log.Fatalf("Failed to load betting queries: %v", err)
		}
	})
	return queries
}
