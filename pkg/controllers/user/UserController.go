package userControllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rafaelmgr12/Mark-URL/pkg/database"
	"github.com/rafaelmgr12/Mark-URL/pkg/models"
	auth "github.com/rafaelmgr12/Mark-URL/pkg/useCase/user"
	"golang.org/x/crypto/bcrypt"
)

type SignupForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ShowUser(c *gin.Context) {
	var user models.User
	database.DB.Find(&user)
	c.JSON(200, user)
}

func Signup(c *gin.Context) {
	var input SignupForm

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usenameAlreadyExist := database.DB.Model(&models.User{}).Where("username = ?", input.Username).First(&models.User{}).Error
	log.Println(usenameAlreadyExist)
	if usenameAlreadyExist == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exist"})
		return
	}

	emailAlreadyExist := database.DB.Model(&models.User{}).Where("email = ?", input.Email).First(&models.User{}).Error
	if emailAlreadyExist == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exist"})
		return
	}

	var user models.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.Password = string(hashedPassword)

	user.Username = input.Username
	user.Password = input.Password
	user.Email = input.Email

	id := uuid.New()
	user.ID = id.String()

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var input LoginForm

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := auth.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
