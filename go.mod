module github.com/echo-marche/presence-api

go 1.13

require (
	github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/jinzhu/gorm v1.9.12
	github.com/mwitkow/go-proto-validators v0.3.0
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.21.0
)

// for realize
replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0
