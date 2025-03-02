module f1betting/betting_system

go 1.23.4

replace f1betting/race_info => ../Race_Info

replace f1betting/dbms => ../Database

replace f1betting/proto => ../proto

require f1betting/race_info v0.0.0

require (
	f1betting/dbms v0.0.0
	f1betting/proto v0.0.0
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.2
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
