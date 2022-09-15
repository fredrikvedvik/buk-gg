import { RouteRecordRaw } from "vue-router";

export default [
    {
        name: "home",
        path: "/",
        component: () => import("@/pages/Home.vue")
    }
] as RouteRecordRaw[];
