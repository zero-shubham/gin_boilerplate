package schemas

type ControllerError struct {
	Err      error
	Messageu string
}

type NotFoundError ControllerError

func (e *NotFoundError) Error() string {
	return e.Err.Error()
}
