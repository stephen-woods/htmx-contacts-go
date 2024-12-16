package model

type Contact struct {
	Errors map[string]string
	First  string
	Last   string
	Email  string
	Phone  string
	Id     int
}
