import Client from "@sanity/client"

const client = new Client({
    token: process.env.SANITY_TOKEN,
    projectId: process.env.SANITY_PROJECT_ID,
})

export type Player = {
    _id: string;
    _type: "player";
    discordId: string;
    location: string;
}

export const getPlayerFromMemberId = async (memberId: string) => {
    const result = await client.fetch("*[_type == 'player' && discordId == $memberId]{_id, _type, discordId, location}", {
        memberId
    }) as Player[]

    if (!result) {
        return null
    }
    
    return result[0]
}