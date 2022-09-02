package api

import (
	"fmt"
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

type Update struct {
	OldUsername string `json:"oldname"`
	NewUsername string `json:"newname"`
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

func DropTable(c *gin.Context) {
	db := db.InitDb()
	res :=db.MustExec(schemas.DropUserTable)
	k,_ := res.RowsAffected()
	c.String(200,"Droped %d",k)

} 


func AddUser(c *gin.Context) {
	var cred User
	db := db.InitDb()

	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured"})
		return
	}

	hashedPass,err := hashpassword.HashPassword(cred.Password)
	if err != nil {
		c.JSON(400,gin.H{"error":"couldnt hash password"})
	}

	db.Exec(schemas.Users)

	Users := `INSERT INTO users(Name,Password) VALUES($1,$2)`

	db.MustExec(Users, cred.Name, hashedPass)

	c.JSON(200,gin.H{
		"success":"Entry created",
	})


}


func CheckPass(c *gin.Context) {
	var cred User

	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured"})
		return
	}

	users := User{}
	db := db.InitDb()

	err := db.Get(&users,"SELECT password FROM users WHERE name=$1",cred.Name)
	if err != nil {
		fmt.Println("error getting data")
		return
	}

	fmt.Println("username from db",users.Name)
	fmt.Println("username from password",users.Password)


	ok := hashpassword.CheckHashPass(cred.Password,users.Password)
	if ok != nil {
		c.JSON(401,gin.H{
			"Wrong Password":"CHeck again",
		})
		
		return
	}

	c.JSON(200,gin.H{
		"success":"verified",
	})

}

func GetUser(c *gin.Context) {
	db := db.InitDb()
	// var cred User
	// if err := c.ShouldBindJSON(&cred); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "error occured"})
	// 	return
	// }
		param := c.Param("name")
	user:= User{}

	err := db.Get(&user,"SELECT name,password FROM users WHERE name=$1",param)
	if err != nil {
		c.String(400,"Couldnt find data")
		return
	}

	c.JSON(200,gin.H{
		"Name":user.Name,
		"Hashed Pass":user.Password,
	})

}

func EditUser(c *gin.Context) {
	db := db.InitDb()
	var cred Update

	oldusername := c.Param("ID")
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured"})
		return
	}

	db.Exec("UPDATE users SET name=$1 WHERE id=$2",cred.NewUsername,oldusername)

	c.JSON(200,gin.H{
		"status":"Updated Successful",
	})


}

func Delete(c *gin.Context) {
	db := db.InitDb()
	id := c.Param("ID")

	db.Exec("DELETE FROM users WHERE id=$1",id)

	c.JSON(200,gin.H{
		"status":"Deleted Successful",
	})


}