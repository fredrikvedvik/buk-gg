package discord

import (
	"encoding/json"
	"fmt"
	"github.com/ansel1/merry/v2"
	"io"
	"net/http"
	"net/url"
	"strings"
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

	req.Header.Set("Authorization", "Bearer "+h.userToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, merry.New("error occurred trying to fetch userToken for code", merry.WithMessage(string(body)), merry.WithHTTPCode(res.StatusCode))
	}

	var user User
	_ = json.Unmarshal(body, &user)

	return &user, nil
}

func (h *Handler) Join(guildID string, userID string, roles []string) error {
	u, err := url.Parse(fmt.Sprintf("https://discord.com/api/v10/guilds/%s/members/%s", guildID, userID))
	if err != nil {
		return err
	}
	reqBody := map[string]any{
		"access_token": h.userToken,
		"roles":        roles,
	}

	marshalled, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("PUT", u.String(), strings.NewReader(string(marshalled)))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bot "+h.client.config.BotToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return merry.New("error occurred trying to fetch userToken for code", merry.WithMessage(string(body)), merry.WithHTTPCode(res.StatusCode))
	}
	return nil
}
