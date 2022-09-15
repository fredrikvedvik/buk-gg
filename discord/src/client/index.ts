import { Client } from "discord.js";

export default new Client({
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
