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
    uuid "github.com/satori/go.uuid"
    "golang.org/x/crypto/bcrypt"
    "google.golang.org/api/googleapi"
    "google.golang.org/api/option"
    "io"
    "net/http"
    "os"
    "strings"
    "time"
)

// This file is for functions that would work with resolvers

var (
    _       = godotenv.Load(".env")
    Envs, _ = godotenv.Read(".env")

    PROJECT_ID        = Envs["GCP_ID"]
    SERVICE_KEY_PATH  = Envs["SERVICE_KEY_PATH"]
    FUNCTION_ENDPOINT = Envs["LOCAL_EMAIL_FUNCTION_ENDPOINT"]
)

func HashPassword(password string) string {
    rawPassword := []byte(password)
    passwordHash, err := bcrypt.GenerateFromPassword(rawPassword, bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Error from hashPassword func")
    }

    return string(passwordHash)
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

func CreateBucket(id int) (string, error) {
    appContext := context.Background()
    // newName  := strings.SplitAfter(name, " ")[0] // removes whitespace && picks name b4 whitespace
    UUID := uuid.NewV4()

    bucketName := fmt.Sprintf("%v-%v", UUID, id)
    client, err := storage.NewClient(appContext, option.WithCredentialsFile(SERVICE_KEY_PATH))
    if err != nil {
        return "", errors.Errorf("some err %v", err)
    }
    fmt.Println(strings.TrimSpace(bucketName))
    // Creates a Bucket instance.
    bucket := client.Bucket(strings.TrimSpace(bucketName))

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
    // f, fileErr := os.Open("go.mod") // to be replaced by passed in file
    //
    // if fileErr != nil {
    //     fmt.Errorf("file not found %v", fileErr)
    // }

    appContext := context.Background()
    client, err := storage.NewClient(appContext, option.WithCredentialsFile(SERVICE_KEY_PATH))

    if err != nil {
       // gCloud uses comma ok idiom not err - nil
    }

    ctx, cancel := context.WithTimeout(appContext, time.Second*50)
    defer cancel()

    wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
    if _, err = io.Copy(wc, file.File); err != nil {
        return errors.Errorf("Error %v", err), nil
    }

    err = wc.Close()

    // i need to sleep the ACL action so the object gets uploaded first
    time.Sleep(1 * time.Second)
    // make file public
    acl := client.Bucket(bucketName).Object(fileName).ACL()
    if err = acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
        return errors.Errorf("Error while setting access %v", err), nil
    }

    if e, ok := err.(*googleapi.Error); ok {
        fmt.Println(e, "comma ok")

        if e.Code == e.Code {
            fmt.Println(e, "trying comma ok")
        }
    }

    return client, nil
}

type details struct {
    email, eventName , eventType string
}

func SendEmail(email string, eventName string, eventType string) (bool, error) {
    reqBody, reqErr := json.Marshal(map[string]string{
        "email": email,
        "type":  eventType,
        "name":  eventName,
    })

    if reqErr != nil {
        return false, errors.Errorf("Error in Marshalling req details %v", reqErr)
    }

    resp, err := http.Post(os.Getenv("FUNCTION_ENDPOINT"), "application/json", bytes.NewBuffer(reqBody))
    if err != nil {
        return false, errors.Errorf("couldn't send email. Error %v" , err)
    }

    fmt.Println(resp)
    defer resp.Body.Close()

    return true, err
}
