package helper

import (
	"exercise/domain"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrUnauthorized:
		return http.StatusNotFound
	case domain.ErrEmailAlreadyExist:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
