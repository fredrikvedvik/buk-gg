package common

import "github.com/google/uuid"

type Settings struct {
	ID         uuid.UUID
	DiscordIDs []string
}
