package main

type UserInterface interface {
	Store(user *User) error
	Get() (user interface{}, err error)
}
