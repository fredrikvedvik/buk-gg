import { Application } from "@/index";

export const getRoleIdsForLocation = (app: Application, location: string) => {
    const roles = app.config.locationRoles.filter(i => i.locations.some(l => l.trim() == location.trim()))

    return roles.map(i => i.roleId)
}