package routes

import (
	"github.com/gin-gonic/gin"
	"rest/controllers"
	"rest/middleware"
)

func SetupBarangRoutes(name string, r *gin.Engine) {
	BarangController := controllers.Brg()
	r.GET(name+"/", middleware.Auth, BarangController.GetAllBarang)
	r.POST(name+"/", middleware.Auth, BarangController.CreateBarang)
}
