package controllers

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"rest/config"
	"rest/models"
	res "rest/service/response"
	"strings"
	"time"
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
	key := os.Getenv("jwt_key")
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), atClaims)
	token, err := sign.SignedString([]byte(key))
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "message": "Successfully authenticated user"})

}

func (repository *LoginController) Register(c *gin.Context) {
	var user models.Users
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	code := c.PostForm("code")

	if email == "" {
		user.UsersEmail = "example@gmail.com"
	}
	if code == "" {
		user.UsersCode = "example@gmail.com"
	}
	if username != "" {
		user.UsersName = username
	}

	hash := md5.Sum([]byte(password))
	var s = hex.EncodeToString(hash[:])
	encryptedpassword := sha1.Sum([]byte(s))
	user.UsersPassword = hex.EncodeToString(encryptedpassword[:])

	//check on user,if not exist create
	row := models.CheckUser(repository.Db, &user, username)
	if row == 0 {
		err := models.CreateUser(repository.Db, &user)
		if err != nil {
			res.Res(c, 400, false, "Registration Failed", nil)
			return
		}
		res.Res(c, http.StatusOK, true, "Registration Success", err)
		return
	}

	res.Res(c, 400, false, "Username is exist!", nil)

}
