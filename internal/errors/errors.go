package errors

import (
	"google.golang.org/grpc"
)

// Define errors
var (
	ErrExceededYourWithdrawalLimit = grpc.Errorf(20001, "You have exceeded your withdrawal limit")
	ErrCannotFindUsers             = grpc.Errorf(20002, "Can not find users")
	ErrAuthenticationFail          = grpc.Errorf(20003, "Please login to access")
)
