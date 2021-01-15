package auth

import (
	"fmt"
	"thinknetica_golang/task_17/pkg/db"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// secretSign ключ подписи JWT
var secretSign = []byte("secret-password")

// Auth объект для авторизации и аутентификации рользователя
type Auth struct{}

// New создает новый объект Auth
func New() *Auth {
	return &Auth{}
}

// Сheck аутентифицирует и авторизует пользователя
func (a *Auth) Сheck(user db.User) (string, error) {
	err := a.authenticate(user)
	if err != nil {
		return "", err
	}

	token, err := a.buildToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Auth) authenticate(user db.User) error {
	password, ok := db.UserPasswords[user.Login]

	if !ok {
		return fmt.Errorf("wrong login")
	}

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		return err
	}

	return nil
}

func (a *Auth) buildToken(user db.User) (string, error) {
	rights, ok := db.AccessRights[user.Login]
	if !ok {
		return "", fmt.Errorf("wrong login")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"rights": rights,
		"usr":    user.Login,
		"nbf":    time.Now().Unix(),
	})
	tokenString, err := token.SignedString(secretSign)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
