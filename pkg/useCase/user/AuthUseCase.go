package auth

import (
	"github.com/rafaelmgr12/Mark-URL/pkg/database"
	"github.com/rafaelmgr12/Mark-URL/pkg/models"
	tokenauth "github.com/rafaelmgr12/Mark-URL/pkg/useCase/token"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(usernam, password string) (string, error) {
	var err error

	u := models.User{}

	err = database.DB.Model(models.User{}).Where("username = ?", usernam).Take(&u).Error
	if err != nil {
		return "", err
	}

	token, err := tokenauth.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
