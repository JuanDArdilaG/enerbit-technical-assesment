package apperrors

func Parse(err error) Error {
	if err == nil {
		return nil
	}
	switch err := err.(type) {
	case Error:
		return err
	default:
		return NewErrorInternalServer(err.Error())
	}
}
