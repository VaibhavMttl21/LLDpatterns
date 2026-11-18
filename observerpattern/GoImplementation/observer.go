package main

import "fmt"

type Observer interface {
	Update(s string)
}

// type UserA struct{}
// type UserB struct{}

// func (u *UserA) Update() {
// 	fmt.Println("notification received from user A")
// }
// func (u *UserB) Update() {
// 	fmt.Println("notification received from user B")
// }

type User struct{
	UserID int
}

func (u * User) Update(msg string){
	fmt.Println("notification received from user with userid",u.UserID ,"with message ", msg )
}