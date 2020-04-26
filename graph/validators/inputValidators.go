package validators

import "github.com/pkg/errors"

var (
    ShortInput = errors.New("input too short")
    NotFound   = errors.New("input not found")
    ErrorUpdating = errors.New("an Error occurred while updating")
)
