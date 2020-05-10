package services

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/echo-marche/presence-api/config"
	pb "github.com/echo-marche/presence-api/proto/pb"
)

func CreateTokenString(req *pb.UserRegistrationRequest) string {
	token := jwt.New(jwt.SigningMethodHS256)
	// Secretで文字列にする. このSecretはサーバだけが知っている
	tokenString, err := token.SignedString([]byte(config.GetEnv("SYSTEM_CODE")))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenString
}
