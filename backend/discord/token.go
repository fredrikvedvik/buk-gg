package discord

type Handler struct {
	client    *Client
	userToken string
}

func (c *Client) NewRequest(userToken string) *Handler {
	return &Handler{
		client:    c,
		userToken: userToken,
	}
}
