package buffalo

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCookies_Get(t *testing.T) {
	r := require.New(t)
	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Cookie", "name=Arthur Dent; answer=42")

	c := Cookies{req, nil}

	v, err := c.Get("name")
	r.NoError(err)
	r.Equal("Arthur Dent", v)

	v, err = c.Get("answer")
	r.NoError(err)
	r.Equal("42", v)

	_, err = c.Get("unknown")
	r.EqualError(err, http.ErrNoCookie.Error())
}

func TestCookies_Set(t *testing.T) {
	r := require.New(t)
	res := httptest.NewRecorder()

	c := Cookies{&http.Request{}, res}

	c.Set("name", "Rob Pike", time.Hour*24)

	h := res.Header().Get("Set-Cookie")
	r.Equal("name=Rob Pike; Max-Age=86400", h)
}

func TestCookies_SetWithExpirationTime(t *testing.T) {
	r := require.New(t)
	res := httptest.NewRecorder()

	c := Cookies{&http.Request{}, res}

	e := time.Date(2017, 7, 29, 19, 28, 45, 0, time.UTC)
	c.SetWithExpirationTime("name", "Rob Pike", e)

	h := res.Header().Get("Set-Cookie")
	r.Equal("name=Rob Pike; Expires=Sat, 29 Jul 2017 19:28:45 GMT", h)
}
