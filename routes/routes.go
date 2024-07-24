package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// lines below for users, not admin 
	// authenticated := server.Group("/")
	// authenticate.Use(middlewares.Authenticate)
	// authenticated.POST("/users/:id")
	server.GET("/users", getEvents)
	server.POST("/login", login)
}


///

//available for users and admin
func getEvents(context *gin.Context) {
	// userID := context.GetInt64("userId")
	
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}

//update is available for admin, but not for users
func updateEvent(context *gin.Context) {
	admin := context.GetBool("admin")

	if admin != true {
		context.JSON(http.StatusUnauthorized, gin.H{"message":"Unauthorized"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "PUT HELLO!"})
}