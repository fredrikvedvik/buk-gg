import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";

import "./index.css";
import auth0 from "./services/auth0";

createApp(App).use(router).use(auth0).mount("#app");
