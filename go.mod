module f1betting

go 1.23.4

replace f1betting/dbms => ./Database

replace f1betting/user_management => ./User_Management

replace f1betting/betting_system => ./Betting_System

replace f1betting/race_info => ./Race_Info

replace f1betting/user_api => ./User_API

require f1betting/dbms v0.0.0

require f1betting/user_management v0.0.0

require (
	f1betting/betting_system v0.0.0
	f1betting/race_info v0.0.0
	f1betting/user_api v0.0.0-00010101000000-000000000000
	github.com/99designs/gqlgen v0.17.64
	github.com/vektah/gqlparser/v2 v2.5.22
)

require (
	github.com/agnivade/levenshtein v1.2.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
