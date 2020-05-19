package validators

import (
    "github.com/pkg/errors"
    "regexp"
)

// Todo this should be called ERROR RESPONSES
var (
    ShortInput           = errors.New("input too short")
    NotFound             = errors.New("User not found")
    LoginError           = errors.New("Password / Email Address invalid")
    ErrorUpdating        = errors.New("an Error occurred while updating")
    QueryError           = errors.New("An error occurred while querying the item")
    ErrorInserting       = errors.New("An error occurred while trying to insert data")
    InvalidEmail         = errors.New("Not a valid Email Address")
    EmailTaken           = errors.New("Email already in use")
    NameTaken            = errors.New("Name already in use")
    TokenGenerationError = errors.New("Error Generating token")
    Unauthorized         = errors.New("Unauthorized access. Login or create account first ")
    ParseToken           = errors.New("Error occurred while parsing token")
)

/*
Todo: This currently doesnt work.
 I want it to break the execution of the function invoking it, so i can clean up the mutations and reuse this block
*/

func LengthChecker(input string, length int) (error, string) {
    if len(input) < length {
        return nil, input
    }

    return nil, input
}

func CheckMail(input string) bool {
    checkMail := regexp.MustCompile(`^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,5})$`)
    result := checkMail.MatchString(input)

    if result == false {
        return false
    }

    return result
}
