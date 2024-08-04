package errors

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func newAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

const (
	ErrBadRequest     = 400
	ErrNotFound       = 404
	ErrInternalServer = 500
	ErrConflict       = 409
)

func NewBadRequest(message string) *AppError {
	return newAppError(ErrBadRequest, message)
}

func NewNotFound(message string) *AppError {
	return newAppError(ErrNotFound, message)
}

func NewInternalServer(message string) *AppError {
	return newAppError(ErrInternalServer, message)
}

func NewConflict(message string) *AppError {
	return newAppError(ErrConflict, message)
}
