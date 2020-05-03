package validators

import "github.com/pkg/errors"
// Todo this should be called ERROR RESPONSES
var (
	ShortInput     = errors.New("input too short")
	NotFound       = errors.New("User not found")
	ErrorUpdating  = errors.New("an Error occurred while updating")
	QueryError     = errors.New("An error occurred while querying the item")
	ErrorInserting = errors.New("An error occurred while trying to insert data")
)

//Todo: Add regex later!
//Checks data from  Mutation input
func LengthChecker( input string, length int)  interface{} {
	if len(input) < length {
		return	ShortInput
	}

	return input
}