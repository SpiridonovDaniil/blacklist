package domain

type Person struct {
	Id       int    `json:"id" example:"1"`
	Phone    string `json:"phone" example:"my_phone"`
	Name     string `json:"name" example:"my_name"`
	Reason   string `json:"reason" example:"my_reason"`
	Time     string `json:"time" example:"16.04.2023"`
	Uploader string `json:"uploader" example:"admin"`
}

type Id struct {
	Id int `json:"id" example:"1"`
}

type Search struct {
	Name  string `json:"name" example:"my_name"`
	Phone string `json:"phone" example:"my_phone"`
}
