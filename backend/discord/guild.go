package discord

type Guild struct {
	Id              string      `json:"id"`
	Name            string      `json:"name"`
	Icon            string      `json:"icon"`
	Description     string      `json:"description"`
	Splash          string      `json:"splash"`
	DiscoverySplash interface{} `json:"discovery_splash"`
	Features        []string    `json:"features"`
	Emojis          []struct {
		Name          string        `json:"name"`
		Roles         []interface{} `json:"roles"`
		Id            string        `json:"id"`
		RequireColons bool          `json:"require_colons"`
		Managed       bool          `json:"managed"`
		Animated      bool          `json:"animated"`
		Available     bool          `json:"available"`
	} `json:"emojis"`
	Stickers          []interface{} `json:"stickers"`
	Banner            string        `json:"banner"`
	OwnerId           string        `json:"owner_id"`
	ApplicationId     interface{}   `json:"application_id"`
	Region            string        `json:"region"`
	AfkChannelId      string        `json:"afk_channel_id"`
	AfkTimeout        int           `json:"afk_timeout"`
	SystemChannelId   string        `json:"system_channel_id"`
	WidgetEnabled     bool          `json:"widget_enabled"`
	WidgetChannelId   interface{}   `json:"widget_channel_id"`
	VerificationLevel int           `json:"verification_level"`
	Roles             []struct {
		Id           string      `json:"id"`
		Name         string      `json:"name"`
		Permissions  string      `json:"permissions"`
		Position     int         `json:"position"`
		Color        int         `json:"color"`
		Hoist        bool        `json:"hoist"`
		Managed      bool        `json:"managed"`
		Mentionable  bool        `json:"mentionable"`
		Icon         *string     `json:"icon"`
		UnicodeEmoji interface{} `json:"unicode_emoji"`
		Flags        int         `json:"flags"`
		Tags         struct {
			PremiumSubscriber interface{} `json:"premium_subscriber"`
		} `json:"tags,omitempty"`
	} `json:"roles"`
	DefaultMessageNotifications int         `json:"default_message_notifications"`
	MfaLevel                    int         `json:"mfa_level"`
	ExplicitContentFilter       int         `json:"explicit_content_filter"`
	MaxPresences                interface{} `json:"max_presences"`
	MaxMembers                  int         `json:"max_members"`
	MaxStageVideoChannelUsers   int         `json:"max_stage_video_channel_users"`
	MaxVideoChannelUsers        int         `json:"max_video_channel_users"`
	VanityUrlCode               string      `json:"vanity_url_code"`
	PremiumTier                 int         `json:"premium_tier"`
	PremiumSubscriptionCount    int         `json:"premium_subscription_count"`
	SystemChannelFlags          int         `json:"system_channel_flags"`
	PreferredLocale             string      `json:"preferred_locale"`
	RulesChannelId              string      `json:"rules_channel_id"`
	PublicUpdatesChannelId      string      `json:"public_updates_channel_id"`
	HubType                     interface{} `json:"hub_type"`
	PremiumProgressBarEnabled   bool        `json:"premium_progress_bar_enabled"`
	Nsfw                        bool        `json:"nsfw"`
	NsfwLevel                   int         `json:"nsfw_level"`
}
