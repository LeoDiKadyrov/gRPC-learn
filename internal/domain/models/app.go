package models

type App struct {
	ID int
	Name string
	Secret string // to sign JWT tokens
}