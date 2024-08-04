package errs

type ServerErr struct {
	Code int
	Msg  string
	Err  error
}

func (e ServerErr) Error() string {
	return e.Msg
}

func SimpleErr(msg string) error {
	return ServerErr{
		Code: 500,
		Msg:  msg,
		Err:  nil,
	}
}

func SimpleErrWithCode(code int, msg string) error {
	return ServerErr{
		Code: code,
		Msg:  msg,
		Err:  nil,
	}
}

func NewServerErr(msg string, err error) error {
	return ServerErr{
		Code: 500,
		Msg:  msg,
		Err:  err,
	}
}

func NewServerErrWithCode(code int, msg string, err error) error {
	return ServerErr{
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}
