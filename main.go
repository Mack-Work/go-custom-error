package main

import (
	customerror "custom-error/error"
	"fmt"
)

func main() {
	err := ReturnErrorNotFound()
	fmt.Println("Not Found: ", customerror.ReturnMessageFromErrorCode(err))

	err = ReturnErrorInvalidInput()
	fmt.Println("Invalid Input: ", customerror.ReturnMessageFromErrorCode(err))

	err = ReturnErrorInternal()
	fmt.Println("Internal Server Error: ", customerror.ReturnMessageFromErrorCode(err))

	err = ReturnErrorUnauthorized()
	fmt.Println("Unauthorized: ", customerror.ReturnMessageFromErrorCode(err))

	err = ReturnErrorForbidden()
	fmt.Println("Forbidden: ", customerror.ReturnMessageFromErrorCode(err))

	err = ReturnErrorBadRequest()
	fmt.Println("Bad Request: ", customerror.ReturnMessageFromErrorCode(err))

	err = ReturnErrorConflict()
	fmt.Println("Conflict: ", customerror.ReturnMessageFromErrorCode(err))

}

// function สำหรับสร้าง error ต่างๆ ตามค่า ErrorCode ที่ระบุ
func ReturnErrorNotFound() error {
	return customerror.NewCustomError("Not Found", customerror.ErrNotFound)
}
func ReturnErrorInvalidInput() error {
	return customerror.NewCustomError("Invalid Input", customerror.ErrInvalidInput)
}
func ReturnErrorInternal() error {
	return customerror.NewCustomError("Internal Server Error", customerror.ErrInternal)
}
func ReturnErrorUnauthorized() error {
	return customerror.NewCustomError("Unauthorized", customerror.ErrUnauthorized)
}
func ReturnErrorForbidden() error {
	return customerror.NewCustomError("Forbidden", customerror.ErrForbidden)
}
func ReturnErrorBadRequest() error {
	return customerror.NewCustomError("Bad Request", customerror.ErrBadRequest)
}
func ReturnErrorConflict() error {
	return customerror.NewCustomError("Conflict", customerror.ErrConflict)
}
