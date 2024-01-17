package go4error

import "fmt"

type GenericError struct {
	Message string
	Code    int
}

func (gerr *GenericError) Error() string {
	return fmt.Sprintf("(CODE: %d) err %s", gerr.Code, gerr.Message)
}
