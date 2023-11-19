package middlewares

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/utils/authutils"
)

func UseAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session := authutils.GetStoreSession(r)

		if _, ok := session.Values["logged"]; !ok {
			w.Header().Set("HX-Redirect", "/login")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next(w, r)
	}

}
