package handler

import (
	"fmt"
	"net/http"

	"github.com/smjn/kubeapp/backend/errors"
)

var NotImplementedErr = fmt.Errorf("Not Implemented")

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	errors.ErrWriter(w, http.StatusNotFound, NotImplementedErr)
})
