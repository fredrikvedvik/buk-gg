package discord

import (
	"encoding/json"
	"github.com/ansel1/merry/v2"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const APIEndpoint = "https://discord.com/api/oauth2/token"

type Config struct {
	ClientID     string
	ClientSecret string
}

type Handler struct {
	config Config
}

func New(config Config) *Handler {
	return &Handler{config: config}
}

type verifyCodeRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
}

func (h *Handler) VerifyCode(code string) (*User, error) {
	u, err := url.Parse(APIEndpoint)
	if err != nil {
		return nil, err
	}

	marshalled, err := json.Marshal(verifyCodeRequest{
		ClientID:     h.config.ClientID,
		ClientSecret: h.config.ClientSecret,
		GrantType:    "authorization_code",
		Code:         code,
	})

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(string(marshalled)))
	if err != nil {
		return nil, err
	}

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

	return nil, nil
}
