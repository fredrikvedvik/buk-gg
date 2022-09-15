package main

import (
	"github.com/fredrikved/buk-gg/discord"
	"github.com/gin-gonic/gin"
)

func discordHandler() gin.HandlerFunc {
	d := discord.New(discord.Config{
		ClientSecret: "",
		ClientID:     "",
	})

	return func(ctx *gin.Context) {
		d.VerifyCode(ctx.Query("code"))
	}
}

func main() {
	r := gin.Default()

	r.POST("discord")

	r.Run(":8002")
}
