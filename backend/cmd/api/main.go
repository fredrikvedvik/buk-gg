package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/fredrikved/buk-gg/auth0"
	"github.com/fredrikved/buk-gg/discord"
	"github.com/gin-gonic/gin"
)

func discordHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Param("token")
		if token == "" {
			ctx.AbortWithStatus(400)
		}

		d := discord.New(token)
		user, _ := d.GetMe()

		ctx.JSON(200, user)
	}
}

func main() {
	ctx := context.Background()
	config := getEnvConfig()
	auth := auth0.New(config.Auth0)

	_, err := firestore.NewClient(ctx, config.Firestore.ProjectID)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(auth.ValidateToken())

	r.POST("discord/:token", discordHandler())

	err = r.Run(":8002")
	if err != nil {
		panic(err)
	}
}
