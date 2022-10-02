import { getOrSet } from "@/storage"
import { Config } from "./types"

export * from "./types"

export const getConfig = async () => {
    return Config.fromObject(await getOrSet("config.json", async () => new Config()))
}