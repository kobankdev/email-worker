package app

import (
	db "email-worker/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func CORSConfig() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}

		context.Next()
	}
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(CORSConfig())

	router.GET("/getEmail/id/:id", server.getEmail)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error:": err.Error()}
}
