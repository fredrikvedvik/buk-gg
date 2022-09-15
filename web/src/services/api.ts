import { api } from "@/config";
import Auth from "./auth";

class API {
    private _token;

    constructor(tokenFactory: () => Promise<string | null>) {
        this._token = tokenFactory;
    }

    private async send(path: string, options: RequestInit) {
        const token = await this._token();
        if (token) {
            options.headers = {
                ...options.headers,
                "Authorization": "Bearer " + token
            }
        }
        const url = new URL(api.url + "/" + path)

        return fetch(url, options).then(r => r.json())
    }

    private async post(path: string, body: any) {
        return this.send(path, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: body ? JSON.stringify(body) : undefined
        })
    }

    public async connectDiscord(token: string) {
        return this.post("discord/" + token, null)
    }
}

export default new API(Auth.getToken)
