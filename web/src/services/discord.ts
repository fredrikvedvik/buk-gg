import api from "./api"
import settings from "./settings"

export const login = () => {
    const url = new URL("https://discord.com/api/oauth2/authorize")
    url.searchParams.append("client_id", import.meta.env.VITE_DISCORD_CLIENT_ID)
    url.searchParams.append("redirect_uri", window.location.origin + "/")
    url.searchParams.append("response_type", "token")
    url.searchParams.append("scope", "identify guilds.join")

    window.location.replace(url)
}

export const completeLogin = async () => {
    const hash = window.location.hash
    const hasAccessToken = hash && hash.includes("access_token")

    if (hasAccessToken) {
        const parts = hash.split("&")

        const values = parts.reduce((a, b) => {
            const ps = b.split("=")
            a[ps[0]] = ps[1]
            return a
        }, {} as {[key: string]: string})

        const accessToken = values.access_token;
        if (accessToken) {
            settings.value = await api.connectDiscord(values.access_token)
            window.location.hash = ""
        }
    }
}

export const removeDiscord = async (id: string) => {
    settings.value = await api.removeDiscord(id)
}

completeLogin()

export default {}