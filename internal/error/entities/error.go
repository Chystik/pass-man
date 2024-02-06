package entities

import (
	"errors"
	"strings"

	"google.golang.org/grpc/codes"
)

type errorCode string

// App error codes.
const (
	ErrNotFound       errorCode = "object not found"
	ErrExists         errorCode = "object already exists in the repository"
	ErrUserCreds      errorCode = "wrong user password"
	ErrAuthClaims     errorCode = "wrong auth claims"
	ErrNotEnoughSpace errorCode = "not enough storage space"

	EDefault errorCode = "internal server error"
)

var codeToGrpcStatus = map[errorCode]codes.Code{
	ErrNotFound:   codes.NotFound,
	ErrExists:     codes.AlreadyExists,
	ErrUserCreds:  codes.Unauthenticated,
	ErrAuthClaims: codes.Unauthenticated,

	EDefault: codes.Internal,
}

type AppError struct {
	// Nested error
	Err error `json:"err"`
	// Error code
	Code errorCode `json:"code"`
	// Error message
	Message string `json:"message"`
	// Executed operation
	Op string `json:"op"`
}

// Error returns the string representation of the error message.
func (e *AppError) Error() string {
	var buf strings.Builder

	if e.Op != "" {
		buf.WriteString(e.Op)
		buf.WriteString(": ")
	}

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			buf.WriteRune('<')
			buf.WriteString(string(e.Code))
			buf.WriteRune('>')
		}
		if e.Code != "" && e.Message != "" {
			buf.WriteRune(' ')
		}
		buf.WriteString(e.Message)
	}

	return buf.String()
}

func ErrorCode(err error) errorCode {
	if err == nil {
		return ""
	}
	target := &AppError{}
	if errors.As(err, &target) {
		if target.Code != "" {
			return target.Code
		}

		if target.Err != nil {
			return ErrorCode(target.Err)
		}
	}

	return EDefault
}

func ErrCodeToGRPC(err error) codes.Code {
	code := ErrorCode(err)
	if v, ok := codeToGrpcStatus[code]; ok {
		return v
	}

	// Default HTTP status for unknown errors
	return codeToGrpcStatus[EDefault]
}
