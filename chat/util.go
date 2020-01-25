package main

import (
	"net/http"
)

func session(w http.ResponseWrite, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = erros.New("Invalid session")
		}
	}
	return
}
