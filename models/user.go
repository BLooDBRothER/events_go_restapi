package models

type User struct {
	Id       string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
