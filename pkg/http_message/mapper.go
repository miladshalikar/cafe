package httpmsg

import (
	"errors"
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
	"net/http"
)

const (
	internalStatus = 500
)

func Error(err error) (message string, code int) {
	var re richerror.RichError
	if !errors.As(err, &re) {
		return err.Error(), http.StatusBadRequest
	}

	msg := re.Message()
	code = mapKindToHTTPStatusCode(re.Kind())

	if code >= internalStatus {
		//todo - log
		//fmt.Println("internal error: ", msg)
		msg = errmsg.ErrorMsgSomethingWentWrong
	}
	return msg, code
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
