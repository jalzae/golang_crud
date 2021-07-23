package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/controllers"
	filter "rest/middleware"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New()
	r.POST("/users", auth, userRepo.CreateUser)
	r.GET("/users", auth, userRepo.GetUsers)
	r.GET("/users/:id", auth, userRepo.GetUser)
	r.PUT("/users/:id", auth, userRepo.UpdateUser)
	r.DELETE("/users/:id", auth, userRepo.DeleteUser)

	BarangController := controllers.Brg()
	r.GET("/Barang", auth, BarangController.GetAllBarang)
	r.POST("/Barang", auth, BarangController.CreateBarang)

	LoginController := controllers.Login()
	r.POST("/Login", LoginController.Login)

	return r

}

//Cors handles cross-domain requests and supports options access
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT,DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//Release all OPTIONS methods
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// process request
		c.Next()
	}
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	key := filter.Getkey()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}

}
