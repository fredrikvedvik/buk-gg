package discord

type Handler struct {
	token string
}

func NewRequest(token string) *Handler {
	return &Handler{token: token}
}
