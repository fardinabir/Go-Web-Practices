package controllers

import (
	"github.com/fardinabir/Go_CRUD_API/model"
	"net/http"
)

type ErrorObj struct {
	Code    string
	Status  int
	Message string
}

func (e *ErrorObj) Error() string {
	return e.Message
}

var (
	ErrUserAlreadyExists        = &ErrorObj{Code: "UA409", Status: http.StatusConflict, Message: "User already exists"}
	ErrUserNotEligible          = &ErrorObj{Code: "UA401", Status: http.StatusConflict, Message: "User is not eligible for registration"}
	ErrNIDAlreadyExists         = &ErrorObj{Code: "UA409", Status: http.StatusConflict, Message: "NID already exists"}
	ErrInvalidDOB               = &ErrorObj{Code: "UA400", Status: http.StatusBadRequest, Message: "Invalid date of birth"}
	ErrSelfieVerificationFailed = &ErrorObj{Code: "UA400", Status: http.StatusBadRequest, Message: "Selfie verification failed"}

	ErrUserCantRequestSelf = &ErrorObj{Code: "UA404", Status: http.StatusNotFound, Message: "User can't request self"}
	ErrUserIDNotFound      = &ErrorObj{Code: "UA404", Status: http.StatusNotFound, Message: "User ID not found"}
	ErrMerchantNotFound    = &ErrorObj{Code: "UA404", Status: http.StatusNotFound, Message: "Merchant not found"}
	ErrAccountNotActive    = &ErrorObj{Code: "UA422", Status: http.StatusUnprocessableEntity, Message: "Account not activated"}
	ErrInvalidToken        = &ErrorObj{Code: "UA422", Status: http.StatusUnprocessableEntity, Message: "Invalid token"}
	ErrTokenExpired        = &ErrorObj{Code: "UA410", Status: http.StatusGone, Message: "Tokens is expired"}
	ErrURINotFound         = &ErrorObj{Code: "UA404", Status: http.StatusNotFound, Message: "URL not found"}
	ErrMethodNotAllowed    = &ErrorObj{Code: "UA405", Status: http.StatusMethodNotAllowed, Message: "Method not allowed"}
	ErrInvalidData         = &ErrorObj{Code: "UA422", Status: http.StatusUnprocessableEntity, Message: "Invalid data"}
	ErrPayloadTooLarge     = &ErrorObj{Code: "UA413", Status: http.StatusRequestEntityTooLarge, Message: "Payload size too large"}
	ErrTooManyRequest      = &ErrorObj{Code: "UA429", Status: http.StatusTooManyRequests, Message: "Too many request"}
	ErrInternalServerError = &ErrorObj{Code: "UA500", Status: http.StatusInternalServerError, Message: "Internal server error"}
	ErrUnauthorizedReq     = &ErrorObj{Code: "UA401", Status: http.StatusUnauthorized, Message: "Unauthorized Access"}
	ErrFailedToCreate      = &ErrorObj{Code: "UA500", Status: http.StatusInternalServerError, Message: "Failed to create the requested"}

	ErrUserNotFound  = &ErrorObj{Code: "UA404", Status: http.StatusNotFound, Message: "User not found"}
	ErrWrongPassword = &ErrorObj{Code: "UA401", Status: http.StatusUnauthorized, Message: "Wrong user credentials"}
)

func (e *ErrorObj) ErrorResponse() *model.Response {
	return &model.Response{Status: e.Status, Body: e}
}
