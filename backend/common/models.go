package common

type Settings struct {
	Discords []*DiscordUser `json:"discords"`
}

type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}

type Config struct {
	GuildID      string   `json:"guildId"`
	MemberRoleID string   `json:"memberRoleId"`
	AdminIDs     []string `json:"adminIds"`
}

type User struct {
	ID      string `json:"id"`
	IsAdmin bool   `json:"isAdmin"`
}
