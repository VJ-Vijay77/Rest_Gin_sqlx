package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTMaker struct {
	secretKey string
}

const minSecretKeySize = 32

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil

}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload,err := NewPayload(username,duration)
	if err != nil {
		return "",err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	return jwtToken.SignedString([]byte(maker.secretKey)) 
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{},error){
		_,ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil,errors.New("token is invalid")
		}
		return []byte(maker.secretKey),nil
	}
	jwtToken,err := jwt.ParseWithClaims(token,&Payload{},keyFunc)
	if err != nil {
		verr,ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner,errors.New("expired token")){
		return nil,errors.New("token expired")
		}
		return nil,errors.New("invalid token")
	}


	payload,ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil,errors.New("invalid token")
	}

	return payload,nil
}
