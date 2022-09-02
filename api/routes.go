package api

import (
	"github.com/VJ-Vijay77/Rest-with-Gin/jwt"
	"github.com/gin-gonic/gin"
)

func Routes(api *gin.Engine) {

	api.POST("/test",Test)
	api.POST("/pass",TestPass)
	api.GET("/drop",DropTable)
	api.POST("/checkpass",CheckPass)
	api.POST("/signin",jwt.Signin)
	api.GET("/welcome",jwt.Welcome)
	api.GET("/refresh",jwt.RefreshJWT)
	api.POST("/adduser",AddUser)
	api.GET("/getuser/:name",GetUser)
	api.PATCH("/edit/:ID",EditUser)
	api.DELETE("/delete/:ID",Delete)

}
