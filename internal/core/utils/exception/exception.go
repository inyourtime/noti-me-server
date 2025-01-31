package exception

import "net/http"

type Err = map[string][]string

type Exception struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Cause   error  `json:"cause"`
	Errors  Err    `json:"errors"`
}

func New(code int, message string, args ...error) *Exception {
	var err error
	if len(args) > 0 {
		err = args[0]
	}
	return &Exception{
		Code:    code,
		Message: message,
		Cause:   err,
		Errors:  make(map[string][]string),
	}
}

func (e *Exception) HasError() bool {
	return len(e.Errors) > 0
}

func (e *Exception) AddError(key, msg string) *Exception {
	e.Errors[key] = append(e.Errors[key], msg)
	return e
}

func Into(err error) *Exception {
	if err == nil {
		return nil
	}
	fail, ok := err.(*Exception)
	if ok {
		return fail
	}
	return New(http.StatusInternalServerError, err.Error(), err)
}

func (e *Exception) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return e.Cause.Error()
}
