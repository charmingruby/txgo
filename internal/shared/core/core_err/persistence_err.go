package core_err

import "fmt"

type PersistenceErr struct {
	Message string `json:"message"`
}

func NewPersistenceErr(originalErr error, action, datasource string) *PersistenceErr {
	return &PersistenceErr{
		Message: fmt.Sprintf("%s persistence error on `%s`: %s", datasource, action, originalErr.Error()),
	}
}

func (e *PersistenceErr) Error() string {
	return e.Message
}
