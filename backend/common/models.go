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
	GuildIDs     []string `json:"guildIds"`
	MemberRoleID string   `json:"memberRoleId"`
	AdminIDs     []string `json:"adminIds"`
}

type Guild struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Joined bool   `json:"joined"`
}

type User struct {
	ID      string `json:"id"`
	IsAdmin bool   `json:"isAdmin"`
}
