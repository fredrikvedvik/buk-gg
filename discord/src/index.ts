import dotenv from "dotenv"
dotenv.config();
import {
    Client
} from "discord.js"
import "./config"
import { Config, getConfig } from "./config";
import { handleVoiceStateUpdate } from "./functions/tempChannels";
import { verifyUser } from "./services/auth";

export class Application {
    public config!: Config
    public client!: Client

    async init() {
        console.log("Initializing")
        this.config = await getConfig();
        this.client = new Client({
            intents: [
                "GuildMembers",
                "GuildMessages",
                "DirectMessageReactions",
                "DirectMessages",
                "GuildVoiceStates",
                "Guilds",
                "GuildPresences"
            ],
        })

        this.client.on("voiceStateUpdate", handleVoiceStateUpdate(this))
        
        await this.client.login(process.env.DISCORD_TOKEN)
    }
}

const service = new Application();

service.init();
