import { auth0 } from "@/config";
import { createAuth0 } from "@auth0/auth0-vue";

const plugin = createAuth0({
    domain: auth0.domain,
    clientId: auth0.clientId,
    cacheLocation: "localstorage",
    authorizationParams: {
        audience: auth0.audience,
        scope: "openid profile email",
        redirect_uri: window.location.origin,
    },
});

export default plugin;

export const useAuth0 = () => {
    return plugin;
};
