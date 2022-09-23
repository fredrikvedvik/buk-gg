package discord

import (
	"encoding/json"
	"github.com/ansel1/merry/v2"
	"io"
	"net/http"
	"net/url"
)

type User struct {
	ID               string      `json:"id"`
	Username         string      `json:"username"`
	Avatar           string      `json:"avatar"`
	AvatarDecoration interface{} `json:"avatar_decoration"`
	Discriminator    string      `json:"discriminator"`
	PublicFlags      int         `json:"public_flags"`
	Flags            int         `json:"flags"`
	Banner           interface{} `json:"banner"`
	BannerColor      string      `json:"banner_color"`
	AccentColor      int         `json:"accent_color"`
	Locale           string      `json:"locale"`
	MfaEnabled       bool        `json:"mfa_enabled"`
	PremiumType      int         `json:"premium_type"`
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
