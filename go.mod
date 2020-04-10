module github.com/echo-marche/presence-api

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.12
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
	google.golang.org/grpc v1.27.0
)

// for realize
replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0
