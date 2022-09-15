package main

import (
	"github.com/fredrikved/buk-gg/auth0"
	"os"
	"strings"
)

type firestoreConfig struct {
	ProjectID string
}

type envConfig struct {
	Auth0     auth0.Config
	Firestore firestoreConfig
}

func getEnvConfig() envConfig {
	audiences := strings.Split(os.Getenv("AUTH0_AUDIENCES"), ",")
	return envConfig{
		Auth0: auth0.Config{
			Audiences: audiences,
			Domain:    os.Getenv("AUTH0_DOMAIN"),
		},
		Firestore: firestoreConfig{
			ProjectID: os.Getenv("FIRESTORE_PROJECT_ID"),
		},
	}
}
