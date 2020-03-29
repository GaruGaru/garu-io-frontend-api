package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		logrus.Error(err)
	}
}
