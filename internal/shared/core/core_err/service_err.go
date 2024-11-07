package core_err

import "fmt"

type ResourceNotFoundErr struct {
	Message string `json:"message"`
}

func NewResourceNotFoundErr(resource string) *ResourceNotFoundErr {
	return &ResourceNotFoundErr{
		Message: fmt.Sprintf("%s not found", resource),
	}
}

func (e *ResourceNotFoundErr) Error() string {
	return e.Message
}

type ResourceAlreadyExistsErr struct {
	Message string `json:"message"`
}

func NewResourceAlreadyExistsErr(resource string) *ResourceAlreadyExistsErr {
	return &ResourceAlreadyExistsErr{
		Message: fmt.Sprintf("%s already exists", resource),
	}
}

func (e *ResourceAlreadyExistsErr) Error() string {
	return e.Message
}