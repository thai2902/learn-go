package model

type Student struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	ClassName   string `json:"class_name"`
	AcademyName string `json:"acedemy_name"`
}
