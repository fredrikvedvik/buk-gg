import { ChannelType, GuildMember, VoiceState } from "discord.js";
import { Application } from "@/index";

export function handleVoiceStateUpdate(app: Application) {
    const createChannelIDs = app.config.tempCreateChannelIDs;

    return async (oldState: VoiceState, newState: VoiceState) => {
        if (newState.channelId === newState.channelId) {
            return
        }
        if (newState.channelId && createChannelIDs.includes(newState.channelId)) {
            await createTemporaryVoiceChannel(app, newState)
        }
        if (oldState.channelId && !createChannelIDs.includes(oldState.channelId)) {
            await handleVoiceChannelLeave(app, oldState)
        }
    }
}

function getChannelNameForMember(app: Application, member: GuildMember) {
    return member.displayName + "'s channel"
}

async function createTemporaryVoiceChannel(app: Application, state: VoiceState) {
    const memberId = state.member?.id;
    if (!memberId) return;

    const parentId = state.channel?.parentId
    if (!parentId) return;

    // Deny the member from creating another channel instantly
    await state.channel.permissionOverwrites.create(memberId, {
        Connect: false,
    })

    // Delete the overwrite after 5 seconds
    setTimeout(async () => {
        if (state.channel) {
            await state.channel.permissionOverwrites.delete(memberId)
        }
    }, 5000)
    
    const newChannel = await state.guild.channels.create({
        type: ChannelType.GuildVoice,
        parent: parentId,
        name: getChannelNameForMember(app, state.member),
        permissionOverwrites: [
            {
                id: state.member,
                allow: [
                    "ManageChannels",
                    "Connect",
                ]
            }
        ]
    })

    await state.member.voice.setChannel(newChannel)
}

async function handleVoiceChannelLeave(app: Application, state: VoiceState) {
    const channelId = state.channelId;
    if (!channelId) {
        return
    }
    if (app.config.permanentChannelIds.includes(channelId)) {
        return
    }

    setTimeout(async () => {
        const channel = await app.client.channels.fetch(channelId)
        if (!channel) {
            return
        }
        if (channel.type === ChannelType.GuildVoice) {
            if (channel.members.size <= 0) {
                await channel.delete("Temporary channel");
            }
        }
    }, 1000)
}
