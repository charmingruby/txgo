package core_err

type EntityErr struct {
	Message string `json:"message"`
}

func NewEntityErr(message string) *EntityErr {
	return &EntityErr{Message: message}
}

func (e *EntityErr) Error() string {
	return e.Message
}
