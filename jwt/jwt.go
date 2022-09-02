package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VJ-Vijay77/Rest-with-Gin/db"
	"github.com/VJ-Vijay77/Rest-with-Gin/hashpassword"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("secret_jwt_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(c *gin.Context) {
	var Creds Credentials

	if err := c.ShouldBindJSON(&Creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured"})
		return
	}

	users := Credentials{}
	db := db.InitDb()

	err := db.Get(&users, "SELECT username,password FROM users WHERE name=$1", Creds.Username)
	if err != nil {
		fmt.Println("error getting data")
		return
	}

	

	ok := hashpassword.CheckHashPass(Creds.Password, users.Password)
	if ok != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Wrong Password": "CHeck again",
		})
		return
	}

	expireTime := time.Now().Add(1 *time.Minute)

	claims := &Claims{
		Username: Creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenString,err := token.SignedString(JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status":"Internal Servor Error",
		})
		return
	}

	http.SetCookie(c.Writer,&http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expireTime,
	})




}
