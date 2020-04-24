package resolvers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// This file is for functions that would work with resolvers

func HashPassword(password string) string {
	rawPassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(rawPassword, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error from hashPassword func")
	}

	convertedToString := string(passwordHash)

	fmt.Println(convertedToString)
	return convertedToString
}
