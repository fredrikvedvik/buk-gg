package discord

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Config struct {
	BotToken     string
	ClientID     string
	ClientSecret string
}

type Client struct {
	config *Config
}

func New(config Config) *Client {
	return &Client{
		config: &config,
	}
}

func (c *Client) GetGuild(ctx context.Context, id string) (*Guild, error) {
	return get[Guild](ctx, c, "guilds/"+id)
}

func (c *Client) RemoveMember(ctx context.Context, guildID string, id string) error {
	_, err := sendDelete[any](ctx, c, "guilds/"+guildID+"/members/"+id)
	return err
}

//func (c *Client) GetMember(ctx context.Context, guild string, id string)

// GetUserAccessToken create an access token for code.
func (c *Client) GetUserAccessToken(code string, redirectUri string) (string, error) {
	u, err := url.Parse("https://discord.com/api/v10/oauth2/token")
	if err != nil {
		return "", err
	}

	q := url.Values{}
	q.Set("client_id", c.config.ClientID)
	q.Set("client_secret", c.config.ClientSecret)
	q.Set("grant_type", "authorization_code")
	q.Set("code", code)
	q.Set("redirect_uri", redirectUri)
	data := q.Encode()

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var t accessToken

	err = json.Unmarshal(body, &t)

	if err != nil {
		return "", err
	}

	return t.AccessToken, nil
}

type accessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}
