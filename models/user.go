package models

type Roles struct {
	Role string `json:"role"`
}

type AuthResponse struct {
	Access   string  `json:"access"`
	Refresh  string  `json:"refresh"`
	FullName string  `json:"full_name"`
	Name     string  `json:"name"`
	Uid      string  `json:"uid"`
	Roles    []Roles `json:"roles"`
	Reason   string  `json:"reason"`
}
