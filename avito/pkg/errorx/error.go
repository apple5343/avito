package errorx

const (
	CodeInternalServerError = 500
	CodeBadRequest          = 400
	CodeNotFound            = 404
	CodeConflict            = 409
	CodeUnauthorized        = 401
)

type Error struct {
	Code int
	Err  error
}

func New(code int, err error) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}

func (e *Error) Unwrap() error {
	return e.Err
}
