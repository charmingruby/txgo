package core_err

import "fmt"

type PersistenceErr struct {
	Message string `json:"message"`
}

func NewPersistenceErr(originalErr error, datasource string) *PersistenceErr {
	return &PersistenceErr{
		Message: fmt.Sprintf("%s persistence error: %s", datasource, originalErr.Error()),
	}
}

func (e *PersistenceErr) Error() string {
	return e.Message
}
