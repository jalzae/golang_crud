package routes

import (
	"github.com/gin-gonic/gin"
	"rest/controllers"
)

func SetupAuthRoutes(name string, r *gin.Engine) {
	LoginController := controllers.Login()
	r.POST(name+"/", LoginController.Login)
	r.POST(name+"/register", LoginController.Register)
}
