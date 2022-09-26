package main

import (
	"github.com/fredrikved/buk-gg/backend/auth0"
	"github.com/fredrikved/buk-gg/backend/discord"
	"os"
	"strings"
)

type firestoreConfig struct {
	ProjectID string
}

type envConfig struct {
	Port      string
	Auth0     auth0.Config
	Firestore firestoreConfig
	Discord   discord.Config
}

func getEnvConfig() envConfig {
	audiences := strings.Split(os.Getenv("AUTH0_AUDIENCES"), ",")
	return envConfig{
		Port: os.Getenv("PORT"),
		Discord: discord.Config{
			ClientID:     os.Getenv("DISCORD_CLIENT_ID"),
			ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
			BotToken:     os.Getenv("DISCORD_BOT_TOKEN"),
		},
		Auth0: auth0.Config{
			Audiences: audiences,
			Domain:    os.Getenv("AUTH0_DOMAIN"),
		},
		Firestore: firestoreConfig{
			ProjectID: os.Getenv("FIRESTORE_PROJECT_ID"),
		},
	}
}
