package resolvers

import (
    "bytes"
    "cloud.google.com/go/storage"
    "context"
    "encoding/json"
    "fmt"
    "github.com/99designs/gqlgen/graphql"
    "github.com/dgrijalva/jwt-go"
    "github.com/joho/godotenv"
    "github.com/pkg/errors"
    "golang.org/x/crypto/bcrypt"
    "google.golang.org/api/option"
    "io"
    "net/http"
    "os"
    "time"
)

// This file is for functions that would work with resolvers

var (
    _       = godotenv.Load(".env")
    Envs, _ = godotenv.Read(".env")

    PROJECT_ID        = Envs["GCP_ID"]
    SERVICE_KEY_PATH  = Envs["SERVICE_KEY_PATH"]
    FUNCTION_ENDPOINT = Envs["EMAIL_FUNCTION_ENDPOINT"]
)

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

func CreateBucket(name string, alias string) (string, error) {
    appContext := context.Background()
    bucketName := name + alias // find a better way to generate names

    client, err := storage.NewClient(appContext, option.WithCredentialsFile(SERVICE_KEY_PATH))
    if err != nil {
        return "", errors.Errorf("some err %v", err)
    }

    // Creates a Bucket instance.
    bucket := client.Bucket(bucketName)

    // Creates the new bucket.
    ctx, cancel := context.WithTimeout(appContext, time.Second*10)
    defer cancel()

    bucketErr := bucket.Create(ctx, PROJECT_ID, nil)
    fmt.Printf("Error %v", bucketErr)

    return bucketName, nil
}

// TODO: Configure gCloud I.A.M to return public file uri

// Accepts {file props} && return uri of uploaded file
func UploadFileToBucket(bucketName string, file graphql.Upload, fileName string) (interface{}, error) {
    f, fileErr := os.Open("go.mod") // to be replaced by passed in file
    appContext := context.Background()

    if fileErr != nil {
        fmt.Errorf("file not found %v", fileErr)
    }

    client, err := storage.NewClient(appContext, option.WithCredentialsFile(SERVICE_KEY_PATH))

    if err != nil {

    }

    ctx, cancel := context.WithTimeout(appContext, time.Second*50)
    defer cancel()

    wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
    if _, err = io.Copy(wc, f); err != nil {
        return errors.Errorf("Error %v", err), nil
    }
    fmt.Println("hello")
    fmt.Println(wc)

    if err := wc.Close(); err != nil {
        return errors.Errorf("Error %v", err), nil
    }

    return client, nil
}

func SendEmail(email string, eventName string, eventType string) (bool, error) {
    sendDetail := fmt.Sprintf("%v?email=%v&type=%v&name=%v", FUNCTION_ENDPOINT, email, eventType, eventName)
    fmt.Println(sendDetail)

    reqBody, reqErr := json.Marshal(map[string]string{
        "email": email,
        "type":  eventType,
        "name":  eventName,
    })

    if reqErr != nil {
        return false, errors.Errorf("Error %v", reqErr)
    }

    resp, err := http.Post(sendDetail, "text/html", bytes.NewBuffer(reqBody))
    if err != nil {
        return false, errors.New("couldn't send email")
    };
    // fmt.Println(resp, "res")
    fmt.Println(resp)
    //resp.Body.Close()

    return true, err
}
