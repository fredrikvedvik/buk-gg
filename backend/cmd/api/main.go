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
	return func(ctx *gin.Context) {
		personID, ok := ctx.Get(auth0.CtxPersonID)
		if !ok {
			ctx.JSON(404, "not found")
			return
		}

		settings, err := store.Settings().Get(ctx, personID.(string))
		if err != nil {
			log.Default().Print("error occurred", err)
			ctx.AbortWithStatus(500)
			return
		}

		if settings == nil {
			settings = &common.Settings{}
		}

		if len(settings.DiscordIDs) >= 2 {
			ctx.JSON(400, "you can not add another discord ID without removing an existing one")
			return
		}

		if settings == nil {
			settings = &common.Settings{}
		}

		token := ctx.Param("token")
		if token == "" {
			ctx.AbortWithStatus(400)
			return
		}

		d := discord.New(token)
		user, _ := d.GetMe()

		if user != nil && !lo.Contains(settings.DiscordIDs, user.ID) {
			settings.DiscordIDs = append(settings.DiscordIDs, user.ID)
			err = store.Settings().Set(ctx, personID.(string), *settings)
			if err != nil {
				log.Default().Print("error occurred", err)
				ctx.AbortWithStatus(500)
				return
			}
		}

		ctx.JSON(200, settings)
	}
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
