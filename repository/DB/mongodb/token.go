package mongodb

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Claims struct {
	Guid string `json:"guid"`
	jwt.StandardClaims
}

var secret = []byte("A&'/}Z57M(2hNg=;LE?")

func CreateRefreshToken(guid string, query func(string, string) error) (string, error) {
	var err error
	var tokenCrypt []byte
	token := make([]byte, 10)
	for i := range token {
		rand.Seed(time.Now().UnixNano())
		token[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	if tokenCrypt, err = bcrypt.GenerateFromPassword(token, 14); err == nil {
		if err = query(string(tokenCrypt), guid); err == nil {
			var tokenStr = base64.StdEncoding.EncodeToString(token)
			return tokenStr, err
		}
	}
	return "", err
}

func GetNewAccessToken(guid string) (string, error) {
	claims := Claims{
		guid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return access.SignedString(secret)
}

func ParseVerifiedAccessToken(token string) (*Claims, error) {
	access, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if access.Valid {
		return access.Claims.(*Claims), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, fmt.Errorf("that's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return access.Claims.(*Claims), fmt.Errorf("Timing is everything")
		}
	}
	return nil, fmt.Errorf("Couldn't handle this token")
}

func RefreshTokenValidate(guid, refresh string) error {
	var err error
	var dbRef *RefreshToken
	var decodeRef []byte
	if dbRef, err = ReadRefreshToken(guid); err == nil {
		if decodeRef, err = base64.StdEncoding.DecodeString(refresh); err == nil {
			if err = bcrypt.CompareHashAndPassword([]byte(dbRef.Refresh), decodeRef); err == nil {
				return nil
			}
		}
	}
	return err
}
