package utils

import "strings"

type ResponseSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObject struct{}

func JsonSuccess(status int, message string, data interface{}) ResponseSuccess {
	return ResponseSuccess{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func JsonError(status int, message string, err string, data interface{}) ResponseError {
	splittedError := strings.Split(err, "/n")
	res := ResponseError{
		Status:  status,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}
