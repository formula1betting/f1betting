module f1betting/betting_system

go 1.23.4

replace f1betting/dbms => ../Database

replace f1betting/user_management => ../User_Management

require f1betting/dbms v0.0.0

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.2
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
