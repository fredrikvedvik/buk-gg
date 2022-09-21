package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/fredrikved/buk-gg/auth0"
	"github.com/fredrikved/buk-gg/common"
	"github.com/fredrikved/buk-gg/database"
	"github.com/fredrikved/buk-gg/discord"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"log"
)

func discordHandler(store *database.Store) gin.HandlerFunc {
	return withSettings(store, func(ctx *gin.Context, userID string, settings *common.Settings) (int, any) {
		token := ctx.Param("token")
		if token == "" {
			return 400, ""
		}

		d := discord.New(token)
		user, _ := d.GetMe()

		if user != nil && !lo.Contains(settings.DiscordIDs, user.ID) {
			settings.DiscordIDs = append(settings.DiscordIDs, user.ID)
			err := store.Settings().Set(ctx, userID, *settings)
			if err != nil {
				log.Default().Print("error occurred", err)
				return 500, "error occurred"
			}
		}

		return 200, settings
	})
}

func withStatusAndResponse(f func(ctx *gin.Context) (int, any)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status, response := f(ctx)
		ctx.JSON(status, response)
	}
}

func withUserID(f func(ctx *gin.Context, userID string) (int, any)) gin.HandlerFunc {
	return withStatusAndResponse(func(ctx *gin.Context) (int, any) {
		personID, ok := ctx.Get(auth0.CtxPersonID)
		if !ok {
			return 404, "not found"
		}

		return f(ctx, personID.(string))
	})
}

func withSettings(store *database.Store, f func(ctx *gin.Context, userID string, settings *common.Settings) (int, any)) gin.HandlerFunc {
	return withUserID(func(ctx *gin.Context, userID string) (int, any) {
		settings, err := store.Settings().Get(ctx, userID)
		if err != nil {
			log.Default().Print("error occurred", err)
			return 500, "error occurred"
		}

		if settings == nil {
			settings = &common.Settings{}
		}

		return f(ctx, userID, settings)
	})
}

func main() {
	ctx := context.Background()
	config := getEnvConfig()
	auth := auth0.New(config.Auth0)

	fs, err := firestore.NewClient(ctx, config.Firestore.ProjectID)
	if err != nil {
		panic(err)
	}

	store := database.New(fs)

	r := gin.Default()
	r.Use(auth.ValidateToken())

	r.POST("discord/:token", discordHandler(store))

	err = r.Run(":8002")
	if err != nil {
		panic(err)
	}
}
