package service

import (
	"admin_panel/model"
	"admin_panel/pkg/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func GetAllUsersFullInfo() (users []model.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for i := range users {
		roles, err := repository.GetAllRolesByUserId(users[i].ID)
		if err != nil {
			return nil, err
		}
		fmt.Println(roles)
		users[i].Roles = roles
	}

	return users, nil
}

func CreateNewUser(user model.User) error {
	user, err := repository.CreateNewUser(user)
	if err != nil {
		log.Println("[service.CreateNewUser]|[repository.CreateNewUser]| error is: ", err.Error())
		return err
	}

	if err := repository.AddRolesToUser(user.ID, user.Roles); err != nil {
		log.Println("[service.CreateNewUser]|[repository.AddRolesToUser]| error is: ", err.Error())
		return err
	}

	return nil
}

func EditUser(role model.User) error {
	return repository.EditUser(role)
}

func DeleteUser(userId int) error {
	return repository.DeleteUser(userId)
}

func AttachRoleToUser(userId, roleId int) error {
	return repository.AttachRoleToUser(userId, roleId)
}

func DetachRoleFromUser(userId, roleId int) error {
	return repository.DetachRoleFromUser(userId, roleId)
}

func GetUserById(userId int) (model.User, error) {
	user, err := repository.GetUserById(userId)
	if err != nil {
		return model.User{}, err
	}

	roles, err := repository.GetAllRolesByUserId(userId)
	if err != nil {
		return model.User{}, err
	}
	user.Roles = roles

	return user, err
}

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	Password string `json:"password"`
}

func GenerateToken(username, password string) (string, error) {
	//user, err := s.repo.GetUser(username, generatePasswordHash(password))
	//if err != nil {
	//	return "", err
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Server",
		},
		password,
	})

	return token.SignedString([]byte(signingKey))
}
