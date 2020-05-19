package resolvers

import (
    "errors"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/joho/godotenv"
    "golang.org/x/crypto/bcrypt"
    "time"
)

// This file is for functions that would work with resolvers

func HashPassword(password string) string {
    rawPassword := []byte(password)
    passwordHash, err := bcrypt.GenerateFromPassword(rawPassword, bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Error from hashPassword func")
    }

    convertedToString := string(passwordHash)
    return convertedToString

}

func GenToken(id string) (string, error) {
    godotenv.Load(".env")
    Envs, _ := godotenv.Read(".env")
    var SECRET = Envs["SECRET_KEY"]

    expiration := time.Now().Add(time.Hour * 24 * 7) // 7 days

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
        ExpiresAt: expiration.Unix(),
        Id:        id, // using user id gets an error. Test with an int
        IssuedAt:  time.Now().Unix(),
        Issuer:    "Oasis",
    })

    accessToken, err := token.SignedString([]byte(SECRET))
    if err != nil {

        // this should return something better
        return accessToken, nil
    }

    return accessToken, errors.New("s")
}

func ComparePassword(userPassword string, password string) interface{} {
    hashedPassword := []byte(password)
    byteUserPassword := []byte(userPassword)

    return bcrypt.CompareHashAndPassword(byteUserPassword, hashedPassword) // returns nil as success
}
