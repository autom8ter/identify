package cookies

import (
	"github.com/autom8ter/identify/util"
	"github.com/gorilla/securecookie"
	"gopkg.in/authboss.v1"
	"log"
	"net/http"
	"time"
)

var CookieStore *securecookie.SecureCookie

func NewSecureCookieStore() {
	tok, _ := util.RandomTokenBytes()
	CookieStore = securecookie.New(tok, nil)
}

type CookieStorer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewCookieStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &CookieStorer{w, r}
}

func (s CookieStorer) Get(key string) (string, bool) {
	cookie, err := s.r.Cookie(key)
	if err != nil {
		return "", false
	}

	var value string
	err = CookieStore.Decode(key, cookie.Value, &value)
	if err != nil {
		return "", false
	}

	return value, true
}

func (s CookieStorer) Put(key, value string) {
	encoded, err := CookieStore.Encode(key, value)
	if err != nil {
		log.Fatal(err)
	}

	cookie := &http.Cookie{
		Expires: time.Now().UTC().AddDate(1, 0, 0),
		Name:    key,
		Value:   encoded,
		Path:    "/",
	}
	http.SetCookie(s.w, cookie)
}

func (s CookieStorer) Del(key string) {
	cookie := &http.Cookie{
		MaxAge: -1,
		Name:   key,
		Path:   "/",
	}
	http.SetCookie(s.w, cookie)
}
