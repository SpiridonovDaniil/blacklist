package domain

type Person struct {
	Id       int    `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Reason   string `json:"reason"`
	Time     string `json:"time"`
	Uploader string `json:"uploader"`
}

type Id struct {
	Id int `json:"id"`
}

type Search struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
