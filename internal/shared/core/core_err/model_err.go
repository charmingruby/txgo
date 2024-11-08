package core_err

type ModelErr struct {
	Message string `json:"message"`
}

func NewModelErr(message string) *ModelErr {
	return &ModelErr{Message: message}
}

func (e *ModelErr) Error() string {
	return e.Message
}
