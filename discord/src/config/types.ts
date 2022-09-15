export class Config {
    public tempCreateChannelIDs: string[] = []
    public permanentChannelIds: string[] = []
    public customChannelNames: string[] = []
    public defaultMemberRoleIDs: string[] = []
    public locationRoles: {
        locations: string[];
        roleId: string;
    }[] = [];

    public static fromObject(object: any) {
        const config = new this()
        if (object) {
            for (const [field, value] of Object.entries(object)) {
                if (!value) continue

                if (typeof ((config as any)[field]) === typeof value) {
                    (config as any)[field] = value
                }
            }
        }
        return config
    }
}