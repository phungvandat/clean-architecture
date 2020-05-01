package errors

type customErr struct {
	message string
	code    int
}

func (ce customErr) Error() string {
	return ce.message
}

func (ce customErr) StatusCode() int {
	return ce.code
}

func newCustomErr(message string, code int) customErr {
	return customErr{
		message: message,
		code:    code,
	}
}
