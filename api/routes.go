package api

import (
	"github.com/VJ-Vijay77/Rest-with-Gin/jwt"
	"github.com/gin-gonic/gin"
)

func Routes(api *gin.Engine) {

	api.POST("/test",Test)
	api.POST("/pass",TestPass)
	api.POST("/adduser",AddUser)
	api.POST("/checkpass",CheckPass)
	api.POST("/signin",jwt.Signin)
}
