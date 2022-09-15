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

        this.client.on("guildMemberAdd", member => {
            if (member.user.bot)
                return;
            
            if (member.pending)
                return;
            
            verifyUser(this, member)
        })

        this.client.on("guildMemberUpdate", (oldMember, newMember) => {
            if (oldMember.user.bot) 
                return;
            
            if (!(oldMember.pending && !newMember.pending)) 
                return;
            
            verifyUser(this, newMember)
        })

        this.client.on("voiceStateUpdate", handleVoiceStateUpdate(this))
        
        await this.client.login(process.env.DISCORD_TOKEN)
        console.log("Logged in as " + this.client.user?.username)
    }
}

const service = new Application();

service.init();
