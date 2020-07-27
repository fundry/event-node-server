package validators

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/softbrewery/gojoi/pkg/joi"
)

var (
	String = joi.String()
	Any    = joi.Any()
)

/*

JOI isn't very useful. Data comes in a specific type and it doesnt check whats in the string.
Example m := "1212". JOI passes (m) as a string causes its wrapped in quotes and casted to a string
==> Check out using REGEXES for that.

*/

func DataValidator(username string) (error, error) {

	// multiple := joi.Slice().Items(
	//     joi.String())

	if err := joi.Validate(username, String); err != nil {
		return nil, errors.New("Failed")
	}

	return nil, errors.New("Failed")
}

func CheckString(s string, i []string) error {

	if check := joi.Validate(s, String); check != nil {
		return errors.Errorf("%v must be a string", s)
	}

	slice := joi.Slice().Items(
		joi.String())

	if err := joi.Validate(i, slice); err != nil {

	}

	return nil
}

// returns true if items in slice are strings
func CheckStrings(i []string) bool {
	slice := joi.Slice().Items(joi.String())

	err := joi.Validate(i, slice)
	fmt.Println(err)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

// breaks a func when value passed is nil
// func RequiredOrBreak(val interface{}) (bool, error) {
//
//     if data := joi.Validate(val, Any); data != nil {
//         return false, errors.Errorf("%v should be more than %v ", val, min)
//     }
//
//     return true, nil
// }

func DataLength(min int, value string, field string) (bool, error) {
	length := joi.String().Min(min)

	if valid := joi.Validate(value, length); valid != nil {
		// fmt.Println(valid)
		return false, errors.Errorf("%v should be more than %v ", field, min)
	}

	return true, nil
}

func BoolRequired(value bool, field string) (bool, error) {
	boolV := joi.Bool().Required()

	if valid := joi.Validate(value, boolV); valid != nil {
		return false, errors.Errorf("%v doesn't contain a Boolean ", field)
	}

	return true, nil
}
