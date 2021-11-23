package model

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	LastName string    `json:"last_name"`
	Login    string    `json:"login"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Status   string    `json:"status"`
	Roles    []RoleDTO `json:"roles,omitempty"`
}

type Role struct {
	Id          int        `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Rights      []RightDTO `json:"rights,omitempty"`
}

type RoleDTO struct {
	Id          int        `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Rights      []RightDTO `json:"omitempty"`
	IsAttached  bool       `json:"is_attached"`
}

type Right struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Section     string `json:"section"`
	Description string `json:"description"`
}

type RightDTO struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Section     string `json:"section"`
	Description string `json:"description"`
	IsAttached  bool   `json:"is_attached"`
}
