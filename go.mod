module github.com/echo-marche/presence-api

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-gorp/gorp v2.2.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/lib/pq v1.1.1 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/mwitkow/go-proto-validators v0.3.0
	github.com/nelsam/hel v2.3.1+incompatible // indirect
	github.com/poy/onpar v0.0.0-20200406201722-06f95a1c68e8 // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.21.0
	gopkg.in/gorp.v2 v2.2.0
)

// for realize
replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0
