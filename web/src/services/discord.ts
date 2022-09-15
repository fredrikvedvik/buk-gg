
export const login = () => {
    const url = new URL("https://discord.com/api/oauth2/authorize")
    url.searchParams.append("client_id", import.meta.env.VITE_DISCORD_CLIENT_ID)
    url.searchParams.append("redirect_uri", window.location.origin + "/")
    url.searchParams.append("response_type", "code")
    url.searchParams.append("scope", "identify")

    window.location.replace(url)
}

export const completeLogin = () => {
    const query = new URLSearchParams(window.location.search)
    const code = query.get("code")
    if (code) {

    }

    // const hash = window.location.hash
    // const hasAccessToken = hash && hash.includes("access_token")

    // if (hasAccessToken) {
    //     const parts = hash.split("&")

    //     const values = parts.reduce((a, b) => {
    //         const ps = b.split("=")
    //         a[ps[0]] = ps[1]
    //         return a
    //     }, {} as {[key: string]: string})

        
    // }
}

export default {}