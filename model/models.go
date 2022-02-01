package model

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	LastName     string    `json:"last_name"`
	Login        string    `json:"login"`
	Email        string    `json:"email"`
	Password     string    `json:"password,omitempty"`
	Status       string    `json:"status"`
	Organization string    `json:"organization"`
	Roles        []RoleDTO `json:"roles,omitempty"`
}

type Role struct {
	ID          int        `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Rights      []RightDTO `json:"rights,omitempty"`
}

type RoleDTO struct {
	ID          int        `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Rights      []RightDTO `json:"omitempty"`
	IsAttached  bool       `json:"is_attached"`
}

type Right struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Section     string `json:"section"`
	Description string `json:"description"`
}

type Cars struct {
	Brand string `json:"brand"`
}

type RightDTO struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Section     string `json:"section"`
	Description string `json:"description"`
	IsAttached  bool   `json:"is_attached"`
}

type SearchContract struct {
	Id             int       `json:"id"`
	Beneficiary    string    `json:"beneficiary"`
	ContractNumber string    `json:"contract_number"`
	ContractType   string    `json:"contract_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Author         string    `json:"author"`
	Price          string    `json:"price"`
}

type ClientBin struct {
	Bin string `json:"bin"`
}

type ContractNumber struct {
	ContractNumber string `json:"contract_number"`
}
