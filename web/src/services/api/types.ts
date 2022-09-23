export type DiscordUser = {
    id: string
    username: string
    discriminator: string
}

export type Settings = {
    discords: DiscordUser[]
}

export type User = {
    id: string
    isAdmin: boolean
}

export type Config = {
    adminIds: string[]
    guildId: string
    memberRoleId: string
}