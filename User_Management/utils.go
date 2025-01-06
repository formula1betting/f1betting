package user_management

import (
	"f1betting/dbms"
)
var userManagementQueries map[string]string

func Init() {
	loadSqlQueries()
}

func loadSqlQueries() {
	userManagementQueries = dbms.LoadUserManagementQueries()
}
