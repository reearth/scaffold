package derror

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
)
