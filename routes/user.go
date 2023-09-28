package routes

import (
	"github.com/gin-gonic/gin"
	"rest/controllers"
	"rest/middleware"
)

func SetupUserRoutes(name string, r *gin.Engine) {
	userRepo := controllers.New()
	r.POST(name+"", middleware.Auth, userRepo.CreateUser)
	r.GET(name+"", middleware.Auth, userRepo.GetUsers)
	r.GET(name+"/:id", middleware.Auth, userRepo.GetUser)
	r.PUT(name+"/:id", middleware.Auth, userRepo.UpdateUser)
	r.DELETE(name+"/:id", middleware.Auth, userRepo.DeleteUser)
}
