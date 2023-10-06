package errors

import "errors"

var (
	ErrorInput          = errors.New("error of input")
	ErrorCase           = errors.New("error of case")
	ErrorHTTP           = errors.New("error http")
	ErrorResponse       = errors.New("error of response")
	ErrorResponseStatus = errors.New("error of response status")
	ErrorNewRequest     = errors.New("error in request")
	ErrorReadBody       = errors.New("error in body response")
	ErrorParseBody      = errors.New("error parse body")
	ErrorAccess         = errors.New("error of access")
)
