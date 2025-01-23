package betting_system

import (
	"f1betting/dbms"
)

var bettingQueries *map[string]string

func Init() {
	loadSqlQueries()
}

func loadSqlQueries() {
	bettingQueries = dbms.GetBettingQueries()
}
