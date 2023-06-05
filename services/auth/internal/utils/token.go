package utils

import (
	"fmt"
	"os"
	"time"

	authpb "github.com/dino9611-grpc-expense-app/grpc-expense-proto/proto"
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	jwt.StandardClaims
	ID int64 `json:"id"`
}

var APPLICATION_NAME = "FINANCE"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodRS256

func ReadRSAKey(dest string) ([]byte, error) {
	// prvKey, err := ioutil.ReadFile("cert/id_rsa")

	key, err := os.ReadFile(dest)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func CreateToken(authpb *authpb.User) error {
	privateKey, err := ReadRSAKey("./cert/secret_rsa")
	if err != nil {
		return fmt.Errorf("read: private key: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	if err != nil {
		return fmt.Errorf("create: parse key: %w", err)
	}
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		ID: authpb.Id,
	}
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(key)

	if err != nil {
		return fmt.Errorf("create: token key: %w", err)
	}
	authpb.Token = signedToken
	return nil
}
