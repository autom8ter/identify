package sessions

import (
	"github.com/autom8ter/identify/util"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"gopkg.in/authboss.v1"
	"log"
)

var SessionCookieName = os.Getenv("SESSION_COOKIE_NAME")

func NewRandomSessionStoreKey() []byte {
	s, _ := util.RandomTokenBytes()
	return s
}

var SessionStore *sessions.CookieStore

func NewSecureSessionCookieStore() {
	SessionStore = sessions.NewCookieStore(NewRandomSessionStoreKey())
}

type SessionStorer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewSessionStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &SessionStorer{w, r}
}

func SetSessionStore(store *sessions.CookieStore) {
	SessionStore = store
}

func GetSessionStore() *sessions.CookieStore {
	return SessionStore
}

func (s SessionStorer) Get(key string) (string, bool) {
	session, err := SessionStore.Get(s.r, SessionCookieName)
	if err != nil {
		log.Println(err.Error())
		return "", false
	}

	strInf, ok := session.Values[key]
	if !ok {
		return "", false
	}

	str, ok := strInf.(string)
	if !ok {
		return "", false
	}

	return str, true
}

func (s SessionStorer) Put(key, value string) {
	session, err := SessionStore.Get(s.r, SessionCookieName)
	if err != nil {
		log.Println(err.Error())
		return
	}

	session.Values[key] = value
	session.Save(s.r, s.w)
}

func (s SessionStorer) Del(key string) {
	session, err := SessionStore.Get(s.r, SessionCookieName)
	if err != nil {
		log.Println(err.Error())
		return
	}

	delete(session.Values, key)
	session.Save(s.r, s.w)
}
