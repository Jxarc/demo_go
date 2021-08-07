package httperror

type HTTPError struct {
	Cause  error  `json:"-"`
	Detail string `json:"details"`
	Status int    `json:"status"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}

	return e.Detail + " : " + e.Cause.Error()
}

func NewHTTPError(err error, status int, detail string) error {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}
