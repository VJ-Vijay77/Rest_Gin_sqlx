package hashpassword

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(pass string) (string,error) {
	HashedPass,err := bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost)
	if err != nil {
		log.Println("Could not hash password")
		return "",err
	}

	return string(HashedPass),nil
}

func CheckHashPass(pass string,hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(pass))
}