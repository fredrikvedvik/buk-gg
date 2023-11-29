import router from "@/router";
import api from "./api";
import Auth from "./auth";
import settings from "./settings";

export const login = () => {
    const url = new URL("https://discord.com/api/oauth2/authorize");
    url.searchParams.append(
        "client_id",
        import.meta.env.VITE_DISCORD_CLIENT_ID
    );
    url.searchParams.append("redirect_uri", window.location.origin + "/");
    url.searchParams.append("response_type", "code");
    url.searchParams.append("scope", "identify guilds.join");

    window.location.replace(url);
};

export const completeLogin = async () => {
    if (!Auth.isAuthenticated().value) {
        return;
    }

    const search = window.location.search.replace("?", "");
    const hasAccessToken = search && search.includes("code");

    if (hasAccessToken) {
        console.log(search);

        const parts = search.split("&");

        const values = parts.reduce((a, b) => {
            const ps = b.split("=");
            a[ps[0]] = ps[1];
            return a;
        }, {} as { [key: string]: string });

        const code = values.code;
        if (code) {
            settings.value = await api.connectDiscord(code);
            router.push({ query: {} });
        }
    }
};

export const removeDiscord = async (id: string) => {
    settings.value = await api.removeDiscord(id);
};

completeLogin();

export default {};
