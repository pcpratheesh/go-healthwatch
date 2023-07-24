package errors

type Error interface {
	Reason() string
	Error() string
}

type err struct {
	msg    string
	reason string
}

func New(msg interface{}, reason string) Error {
	var errorMessage string

	switch val := msg.(type) {
	case error:
		errorMessage = val.Error()
	case string:
		errorMessage = val
	}

	return &err{
		msg:    errorMessage,
		reason: reason,
	}
}

func (e *err) Error() string {
	return e.msg
}

func (e *err) Reason() string {
	return e.reason
}
