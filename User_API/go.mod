module f1betting/user_api

replace f1betting/user_management => ../User_Management

replace f1betting/betting_system => ../Betting_System

replace f1betting/dbms => ../Database

replace f1betting/race_info => ../Race_Info

require (
	f1betting/betting_system v0.0.0
	github.com/99designs/gqlgen v0.17.64
	github.com/jackc/pgx/v5 v5.7.2
	github.com/vektah/gqlparser/v2 v2.5.22
)

require (
	f1betting/dbms v0.0.0 // indirect
	f1betting/race_info v0.0.0 // indirect
	github.com/agnivade/levenshtein v1.2.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/urfave/cli/v2 v2.27.5 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

go 1.23.4
