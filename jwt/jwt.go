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
	Username string `json:"username" db:"name"`
	Password string `json:"password" db:"password"`
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

	err := db.Get(&users, "SELECT * FROM users WHERE name=$1", Creds.Username)
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

	expireTime := time.Now().Add(9 * time.Minute)
	JwtexpireTime := time.Now().Add(25 * time.Second)

	claims := &Claims{
		Username: Creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: JwtexpireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Internal Servor Error",
		})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expireTime,
	})

	c.String(200,"Logged in")

}

func Welcome(c *gin.Context) {
	ck, err := c.Request.Cookie("token")
	if err != nil {

		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "UnAuthorized(cookie)",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request from cookie"})
		return
	}

	tokenString := ck.Value

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "UnAuthorized(parse with cliaM)",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request(PARSE WITH CLAIMS)"})
		return
	}

	if !token.Valid {
		c.JSON(401, gin.H{
			"status": "Status UnAuthorized (VALID)",
		})
		return
	}

	c.String(200, "Welcome ,jwt authorized(VALID)")

}


func RefreshJWT (c *gin.Context) {
	ck, err := c.Request.Cookie("token")
	if err != nil {

		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "UnAuthorized(cookie)",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request from cookie"})
		return
	}

	tokenString := ck.Value

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if !token.Valid {
		c.JSON(401, gin.H{
			"status": "Status Token not valid (VALID)",
		})
		
	}


	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "UnAuthorized(parse with cliaM)",
			})
			
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request(PARSE WITH CLAIMS)"})
		
	}


	if time.Unix(claims.ExpiresAt,0).Sub(time.Now()) > 35*time.Second {
		c.JSON(400,gin.H{"staus":"Bad Request (time 35 + seconds)"})
		return
	}

	JwtExpireTime := time.Now().Add(25 *time.Second)
	ExpireTime := time.Now().Add(5 *time.Minute)
	claims.ExpiresAt = JwtExpireTime.Unix()
	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	newTokenString,err := newtoken.SignedString(JwtKey)
	if err != nil {
		c.JSON(500,gin.H{"status":"Internal Server Error(new token string)"})
		return
	}

	http.SetCookie(c.Writer,&http.Cookie{
		Name: "token",
		Value: newTokenString,
		Expires: ExpireTime,
	})

	c.String(200,"New Token Generated")

}