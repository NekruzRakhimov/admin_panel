package models

type Roles struct {
	Role string `json:"role"`
}


type AuthResponse struct {
	Access string `json:"access"`
	Refresh string `json:"refresh"`
	Name  string `json:"name"`
	Uid   string `json:"uid"`
	Roles []Roles `json:"roles"`
}