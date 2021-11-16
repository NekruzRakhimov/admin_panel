package service

import (
	"admin_panel/models"
	"admin_panel/pkg/repository"
	"fmt"
	"log"
)

func GetAllUsersFullInfo() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for i := range users {
		roles, err := repository.GetAllRolesByUserId(users[i].Id)
		if err != nil {
			return nil, err
		}
		fmt.Println(roles)
		users[i].Roles = roles
	}

	return users, nil
}

func CreateNewUser(user models.User) error {
	user, err := repository.CreateNewUser(user)
	if err != nil {
		log.Println("[service.CreateNewUser]|[repository.CreateNewUser]| error is: ", err.Error())
		return err
	}

	if err := repository.AddRolesToUser(user.Id, user.Roles); err != nil {
		log.Println("[service.CreateNewUser]|[repository.AddRolesToUser]| error is: ", err.Error())
		return err
	}

	return nil
}

func EditUser(role models.User) error {
	return repository.EditUser(role)
}

func DeleteUser(userId int) error {
	return repository.DeleteUser(userId)
}

func AttachRoleToUser (userId, roleId int) error {
	return repository.AttachRoleToUser(userId, roleId)
}

func DetachRoleFromUser (userId, roleId int) error {
	return repository.DetachRoleFromUser(userId, roleId)
}