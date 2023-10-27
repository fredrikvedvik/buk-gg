import { RouteRecordRaw } from "vue-router";

export default [
    {
        name: "home",
        path: "/",
        component: () => import("@/pages/Home.vue"),
        children: [
            {
                name: "discord",
                path: "",
                component: () => import("@/pages/Discord.vue"),
            },
            {
                name: "rules",
                path: "rules",
                component: () => import("@/pages/Rules.vue"),
            },
            {
                name: "games",
                path: "games",
                component: () => import("@/pages/Games.vue"),
            },
        ],
    },
] as RouteRecordRaw[];
