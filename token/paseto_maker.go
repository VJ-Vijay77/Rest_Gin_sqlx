package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)


type PasetoMaker struct {
	paseto *paseto.V2
	symmetricalKey []byte
}

func NewPasetoMaker(symmetricalKey string) (Maker,error) {
	if len(symmetricalKey) != chacha20poly1305.KeySize {
		return nil,fmt.Errorf("invalid key size: must be exactly %d characters",chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		symmetricalKey: []byte(symmetricalKey),
	}
	return maker, nil
}


func(maker *PasetoMaker) CreateToken(username string,duration time.Duration) (string,error) {
	payload,err := NewPayload(username,duration)
	if err != nil {
		return "",err
	}

	return maker.paseto.Encrypt(maker.symmetricalKey,payload,nil)
}


func(maker *PasetoMaker) VerifyToken(token string) (*Payload,error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token,maker.symmetricalKey,payload,nil)
	if err != nil {
		return nil,errors.New("invalid token")
	}

	err = payload.Valid()
	if err != nil {
		return nil,err
	}

	return payload,nil

}