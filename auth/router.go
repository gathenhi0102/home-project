package auth

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func AuthRouter(router *httprouter.Router) {
	router.GET("/auth/login", Hello)
}
