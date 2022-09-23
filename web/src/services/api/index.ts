import { api } from "@/config";
import Auth from "../auth";
import { Config, Settings, User } from "./types";
export * from "./types"

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

    private async post<T>(path: string, body: any) {
        return this.send(path, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: body ? JSON.stringify(body) : undefined
        }) as Promise<T>
    }

    private async get<T>(path: string) {
        return this.send(path, {
            method: "GET",
        }) as Promise<T>
    }

    private async delete<T>(path: string) {
        return this.send(path, {
            method: "DELETE"
        }) as Promise<T>
    }

    public connectDiscord(token: string) {
        return this.post<Settings>("discord/" + token, null)
    }

    public removeDiscord(id: string) {
        return this.delete<Settings>("discord/" + id)
    }

    public getSettings() {
        return this.get<Settings>("settings")
    }

    public getUser() {
        return this.get<User>("user")
    }

    public getConfig() {
        return this.get<Config>("config")
    }

    public setConfig(config: Config) {
        return this.post<Config>("config", config)
    }
}

export default new API(Auth.getToken)
