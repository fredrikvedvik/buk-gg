package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/fredrikved/buk-gg/auth0"
	"github.com/fredrikved/buk-gg/common"
	"github.com/fredrikved/buk-gg/database"
	"github.com/fredrikved/buk-gg/discord"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"log"
)

type handlers struct {
	dc     *discord.Client
	config *common.Config
	store  *database.Store
}

func (h *handlers) discordHandler() gin.HandlerFunc {

	return h.withSettings(func(ctx *gin.Context, userID string, settings *common.Settings) (int, any) {
		code := ctx.Param("code")
		if code == "" {
			return 400, ""
		}

		redirectUri := ctx.Query("redirectUri")

		token, err := h.dc.GetUserAccessToken(code, redirectUri)

		d := h.dc.NewRequest(token)
		user, _ := d.GetMe()

		if user == nil {
			return 200, settings
		}

		stored, found := lo.Find(settings.Discords, func(d *common.DiscordUser) bool {
			return d.ID == user.ID
		})
		if stored == nil {
			stored = &common.DiscordUser{
				ID: user.ID,
			}
		}
		stored.Username = user.Username
		stored.Discriminator = user.Discriminator
		if !found {
			if len(settings.Discords) >= 2 {
				return 400, settings
			}
			settings.Discords = append(settings.Discords, stored)
		}
		err = h.store.Settings().Set(ctx, userID, *settings)
		if err != nil {
			log.Default().Print("error occurred", err)
			return 500, "error occurred"
		}

		err = d.Join(h.config.GuildID, user.ID, []string{h.config.MemberRoleID})
		if err != nil {
			log.Default().Print("error occurred", err)
			return 500, "error occurred"
		}

		return 200, settings
	})
}

func (h *handlers) deleteDiscordHandler() gin.HandlerFunc {
	return h.withSettings(func(ctx *gin.Context, userID string, settings *common.Settings) (int, any) {
		id := ctx.Param("id")
		if id == "" {
			return 400, ""
		}

		if lo.SomeBy(settings.Discords, func(d *common.DiscordUser) bool {
			return d.ID == id
		}) {
			settings.Discords = lo.Filter(settings.Discords, func(d *common.DiscordUser, _ int) bool {
				return d.ID != id
			})
			err := h.store.Settings().Set(ctx, userID, *settings)
			if err != nil {
				return 500, err
			}
		} else {
			return 200, "id not found"
		}

		return 200, settings
	})
}

func (h *handlers) settingsHandler() gin.HandlerFunc {
	return h.withSettings(func(ctx *gin.Context, userID string, settings *common.Settings) (int, any) {
		return 200, settings
	})
}

func (h *handlers) getConfigHandler() gin.HandlerFunc {
	return h.withUser(func(ctx *gin.Context, user *common.User) (int, any) {
		if !lo.Contains(h.config.AdminIDs, user.ID) {
			return 403, "no access"
		}
		config, err := h.store.Config().Get(ctx, "general")
		if err != nil {
			return 500, err
		}
		return 200, config
	})
}

func (h *handlers) updateConfigHandler() gin.HandlerFunc {
	return h.withUser(func(ctx *gin.Context, user *common.User) (int, any) {
		if !lo.Contains(h.config.AdminIDs, user.ID) {
			return 403, "no access"
		}
		var config common.Config

		err := ctx.BindJSON(&config)
		if err != nil {
			return 400, err
		}

		err = h.store.Config().Set(ctx, "general", config)
		if err != nil {
			return 500, err
		}
		return 200, config
	})
}

func (h *handlers) currentUserHandler() gin.HandlerFunc {
	return h.withUser(func(ctx *gin.Context, user *common.User) (int, any) {
		return 200, user
	})
}

func (h *handlers) withStatusAndResponse(f func(ctx *gin.Context) (int, any)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status, response := f(ctx)
		ctx.JSON(status, response)
	}
}

func (h *handlers) withUser(f func(ctx *gin.Context, user *common.User) (int, any)) gin.HandlerFunc {
	return h.withStatusAndResponse(func(ctx *gin.Context) (int, any) {
		personID, ok := ctx.Get(auth0.CtxPersonID)
		if !ok {
			return 404, "not found"
		}

		userID := personID.(string)

		return f(ctx, &common.User{
			ID:      userID,
			IsAdmin: lo.Contains(h.config.AdminIDs, userID),
		})
	})
}

func (h *handlers) withSettings(f func(ctx *gin.Context, userID string, settings *common.Settings) (int, any)) gin.HandlerFunc {
	return h.withUser(func(ctx *gin.Context, user *common.User) (int, any) {
		settings, err := h.store.Settings().Get(ctx, user.ID)
		if err != nil {
			log.Default().Print("error occurred", err)
			return 500, "error occurred"
		}

		if settings == nil {
			settings = &common.Settings{}
		}

		return f(ctx, user.ID, settings)
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

	cfg, err := store.Config().Get(ctx, "general")
	if err != nil {
		panic(err)
	}
	if cfg == nil {
		cfg = &common.Config{}
		err = store.Config().Set(ctx, "general", *cfg)
		if err != nil {
			panic(err)
		}
	}

	h := &handlers{
		dc:     discord.New(config.Discord),
		config: cfg,
		store:  store,
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Authorization", "Content-Type"},
	}))
	r.Use(auth.ValidateToken())
	r.GET("config", h.getConfigHandler())
	r.POST("config", h.updateConfigHandler())

	r.GET("user", h.currentUserHandler())
	r.GET("settings", h.settingsHandler())

	r.POST("discord/:code", h.discordHandler())
	r.DELETE("discord/:id", h.deleteDiscordHandler())

	err = r.Run(":8002")
	if err != nil {
		panic(err)
	}
}
