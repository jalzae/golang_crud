package controllers

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"rest/config"
	"rest/models"
	"strings"
	"time"
	filter "rest/middleware"
)

type LoginController struct {
	Db *gorm.DB
}

func Login() *LoginController {
	db := config.InitDb()
	return &LoginController{Db: db}
}

func (repository *LoginController) Login(c *gin.Context) {

	var user models.Users
	username := c.PostForm("username")
	password := c.PostForm("password")

	hash := md5.Sum([]byte(password))
	var s = hex.EncodeToString(hash[:])
	encryptedpassword := sha1.Sum([]byte(s))
	var realpassword = hex.EncodeToString(encryptedpassword[:])
	row := models.LoginUser(repository.Db, &user, username, realpassword)
	// Validate form input

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	if row != 1 {
		c.AbortWithStatusJSON(400, gin.H{"error": "User salah"})
		return
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
key := filter.Getkey()
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), atClaims)
	token, err := sign.SignedString([]byte(key))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "message": "Successfully authenticated user"})

}
