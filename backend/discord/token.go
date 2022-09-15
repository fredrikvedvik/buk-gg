package discord

import (
	"encoding/json"
	"github.com/ansel1/merry/v2"
	"io"
	"net/http"
	"net/url"
)

type Handler struct {
	token string
}

func New(token string) *Handler {
	return &Handler{token: token}
}

func (h *Handler) GetMe() (*User, error) {
	u, err := url.Parse("https://discord.com/api/v10/users/@me")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+h.token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, merry.New("error occured trying to fetch token for code", merry.WithMessage(string(body)), merry.WithHTTPCode(res.StatusCode))
	}

	var user User
	_ = json.Unmarshal(body, &user)

	return &user, nil
}
