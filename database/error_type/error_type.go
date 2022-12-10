package error_type

import "fmt"

var (
	ErrNotFound     = fmt.Errorf("not Found")
	ErrPostNotExist = fmt.Errorf("post not exist")
)
