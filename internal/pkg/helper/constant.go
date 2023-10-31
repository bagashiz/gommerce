package helper

import "errors"

// Constants for message in HTTP response
const (
	FAILEDPOSTDATA    = "Failed to POST data"
	FAILEDGETDATA     = "Failed to GET data"
	FAILEDPUTDATA     = "Failed to PUT data"
	FAILEDDELETEDATA  = "Failed to DELETE data"
	SUCCEEDPOSTDATA   = "Succeed to POST data"
	SUCCEEDGETDATA    = "Succeed to GET data"
	SUCCEEDPUTDATA    = "Succeed to PUT data"
	SUCCEEDDELETEDATA = "Succeed to DELETE data"
)

// Error messages
var (
	// ErrUnsupportedDriver is an error message for unsupported database driver
	ErrUnsupportedDriver = errors.New("unsupported database driver")
	// ErrUnsupportedTokenType is an error message for unsupported token type
	ErrUnsupportedTokenType = errors.New("unsupported token type")
	// ErrDataNotFound is an error message for requested data not found
	ErrDataNotFound = errors.New("data not found")
	// ErrDataAlreadyExists is an error message for unique key constraint violation
	ErrDataAlreadyExists = errors.New("data already exists")
	// ErrUnauthorized is an error message for unauthorized access
	ErrUnauthorized = errors.New("unauthorized")
	// ErrForbidden is an error message for forbidden access
	ErrForbidden = errors.New("forbidden")
	// ErrExpiredToken is an error message for expired token
	ErrExpiredToken = errors.New("token has expired")
	// ErrInvalidToken is an error message for invalid token
	ErrInvalidToken = errors.New("token is invalid")
)
