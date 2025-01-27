package user_management

import (
	"f1betting/dbms"
)

var UserManagementQueries *map[string]string

func Init() {
	loadSqlQueries()
}

func loadSqlQueries() {
	UserManagementQueries = dbms.GetUserManagementQueries()
}
