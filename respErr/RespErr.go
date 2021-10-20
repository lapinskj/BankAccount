package respErr

type RespErr struct {
	Err        error
	StatusCode int
}

func (respErr RespErr) Error() string {
	return respErr.Err.Error()
}

func (respErr RespErr) Status() int {
	return respErr.StatusCode
}
