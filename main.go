package main

import (
	"fmt"
)

type User struct {
	firstname string
	lastname  string
}

func (userReference User) getFullName() string {
	return userReference.firstname + " " + userReference.lastname
}
func main() {
	firstname := "Gratsiya"
	lastname := "Kalinina"
	user := User{
		firstname: firstname,
		lastname:  lastname,
	}
	fmt.Print(user)
}
