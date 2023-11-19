package authutils

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func Initialize() {
	sessionKey := []byte(os.Getenv("SESSION_KEY"))
	if len(sessionKey) == 0 {
		log.Fatal("SESSION_KEY environment variable is not set")
	}

	store = sessions.NewCookieStore(sessionKey)
}

func GetStoreSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "session-auth")
	return session
}

func SetUserSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "session-auth")
	session.Values["logged"] = true
	session.Options.MaxAge = 60 * 60 * 24
	err := session.Save(r, w)
	return err
}

func ExpireSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "session-auth")
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	return err
}
