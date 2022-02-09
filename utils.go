package wework

func NewResult(errCode int, errMsg string, err error) *Result {
	return &Result{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Error:   err,
	}
}
