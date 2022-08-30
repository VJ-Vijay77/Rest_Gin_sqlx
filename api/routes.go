package api

import (
	"github.com/gin-gonic/gin"
)

func Routes(api *gin.Engine) {

	api.POST("/test",Test)
	api.POST("/pass",TestPass)
}
