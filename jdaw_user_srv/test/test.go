package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	s := string(hasedPassword)
	if err == nil {
		fmt.Println(s)
	}
}
