package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ansel1/merry/v2"
	"io"
	"net/http"
	"net/url"
)

func get[T any](ctx context.Context, client *Client, path string) (*T, error) {
	u, err := url.Parse(fmt.Sprintf("https://discord.com/api/v10/%s", path))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bot "+client.config.BotToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resultString, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, merry.New(string(resultString))
	}

	var result T
	err = json.Unmarshal(resultString, &result)
	return &result, err
}

func sendDelete[T any](ctx context.Context, client *Client, path string) (*T, error) {
	u, err := url.Parse(fmt.Sprintf("https://discord.com/api/v10/%s", path))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bot "+client.config.BotToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resultString, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, merry.New(string(resultString))
	}

	var result T
	err = json.Unmarshal(resultString, &result)
	return &result, err
}
