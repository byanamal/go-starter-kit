package constants

import "errors"

var (
	INTERNAL_SERVER_ERROR = "Internal Server Error"
	INVALID_REQUEST       = "Invalid Request"
	INVALID_FORMAT_UUID   = "Invalid UUID Format"
	ID_IS_REQUIRED        = "ID Is Required"
	INVALID_CREDENTIAL    = "Email Or Password Wrong"
	UNAUTHORIZED          = "Unauthorized"
	FORBIDDEN             = "Forbidden"
)

var ErrUserNotFound = errors.New("User Not Found")
var ErrInvalidCredential = errors.New("Email Or Password Wrong")
