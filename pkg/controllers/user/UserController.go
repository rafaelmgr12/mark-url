package userControllers

import (
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

func Signup(c *gin.Context) {
	var input SignupForm
	var err error

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	err = database.DB.Model(models.User{}).Where("username = ?", input.Username).Take(&u).Error
	println(err)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already taken."})
		return
	}

	err = database.DB.Model(models.User{}).Where("email = ?", input.Email).Take(&u).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already taken."})
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
