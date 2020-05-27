package model

// Todo: Move this into same file with tester route
type BetaTester struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    DateApplied string `json:"dateApplied"`
}
