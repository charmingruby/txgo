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

type InvalidFundsErr struct {
	Message string `json:"message"`
}

func NewInvalidFundsErr(missingPoints int) *InvalidFundsErr {
	return &InvalidFundsErr{
		Message: fmt.Sprintf("invalid funds, missing %d points", missingPoints),
	}
}

func (e *InvalidFundsErr) Error() string {
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

type ForbiddenActionErr struct {
	Message string `json:"message"`
}

func NewForbiddenActionErr(reason string) *ForbiddenActionErr {
	return &ForbiddenActionErr{
		Message: reason,
	}
}

func (e *ForbiddenActionErr) Error() string {
	return e.Message
}
