import { getOrSet } from "src/storage"
import { Config } from "./types"
import "./web"

export * from "./types"

export const getConfig = async () => {
    return Config.fromObject(await getOrSet("config.json", async () => new Config()))
}