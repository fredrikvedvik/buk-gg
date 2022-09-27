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
    guildIds: string[]
    memberRoleId: string
}

export type Guild = {
    id: string
    name: string
    icon: string
}