import { GuildMember } from "discord.js";
import { Application } from "@/index";
import {getPlayerFromMemberId} from "@/sanity"
import { getRoleIdsForLocation } from "@/roles";

export const verifyUser = async (app: Application, member: GuildMember) => {
    const player = await getPlayerFromMemberId(member.id)

    if (!player) {
        return
    }

    const guild = member.guild;

    for (const roleId of app.config.defaultMemberRoleIDs) {
        const role = await guild.roles.fetch(roleId)

        if (role) {
            await member.roles.add(role)
        } else {
            console.error("Missing role by id: " + roleId)
        }
    }

    const locationRoleIds = getRoleIdsForLocation(app, player.location)

    for (const roleId of locationRoleIds) {
        const role = await guild.roles.fetch(roleId)

        if (role) {
            await member.roles.add(role)
        } else {
            console.error("Missing role by id: " + roleId)
        }
    }

    await member.send("Welcome to the official BUK Discord server!")
}