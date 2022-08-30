package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)
type Server struct {
	db *sqlx.DB
	router *gin.Engine
}

func ReturnServer(db sqlx.DB) *Server{
	return &Server{
		db: &db,
	}
}

func NewServer(store *sqlx.DB) *gin.Engine {

	server := &Server{db: store}
	router := gin.Default()
	Routes(router)
		
	server.router = router

	return router
}

