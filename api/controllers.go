package api

import (
	"net/http"

	"github.com/VJ-Vijay77/Rest-with-Gin/db"
	"github.com/VJ-Vijay77/Rest-with-Gin/hashpassword"
	"github.com/VJ-Vijay77/Rest-with-Gin/schemas"
	"github.com/gin-gonic/gin"
)

type Details struct {
	Name   string `json:"name"`
	Job    string `json:"job"`
	Salary int64  `json:"salary"`
	Age    int64  `json:"age"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"pass"`
}

type Password struct {
	Password string `json:"pass"`
}

func Test(c *gin.Context) {
	var req Details
	db := db.InitDb()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured"})
		return
	}

	db.Exec(schemas.Place)
	db.Exec(schemas.AlterAge)
	//  if err != nil {
	// 	log.Println("error getting schema")
	//  }
	Name := `INSERT INTO details(Name,Job,Salary,Age) VALUES($1,$2,$3,$4)`
	db.MustExec(Name, req.Name, req.Job, req.Salary, req.Age)

	c.JSON(200, gin.H{"Success": "Successfully Created"})

}

func TestPass(c *gin.Context) {
	var req Password

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured"})
		return
	}

	hashed, err := hashpassword.HashPassword(req.Password)
	if err != nil {
		c.JSON(401, gin.H{"e": "Error getting hash"})
		return
	}
	var result string
	check := hashpassword.CheckHashPass(req.Password, hashed)
	if check == nil {
		result = "Verified"
	} else {
		result = "Wrong one"
	}
	c.JSON(200, gin.H{
		"Hashed Password":         hashed,
		"Checking Passed if Same": result,
	})
}

func AddUser(c *gin.Context) {
	var cred User

	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured"})
		return
	}
}
