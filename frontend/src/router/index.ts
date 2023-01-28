import { createRouter, createWebHistory, NavigationGuard, Router, RouteRecordRaw } from "vue-router";
import { nextTick } from "vue"

declare module "vue-router" {
    interface RouteMeta {
        requiresAuth?: boolean;
        requiresRole?: string[] | string;
        cidOverridesRole?: boolean;
    }
}

const routes = [
    {
        path: "/",
        name: "Home",
        component: () => import("@/views/Home.vue"),
    },
    {
        path: "/search",
        name: "Search",
        component: () => import("@/views/Search.vue"),
    },
    {
        path: "/not-implemented",
        name: "NotImplemented",
        component: () => import("@/views/NotImplementedPage.vue"),
    },
    {
        path: "/:pathMatch(.*)*",
        name: "NotFound",
        component: () => import("@/views/errors/ErrorNotFound.vue"),
    },
];

const buildRouter = (): Router => {
    return createRouter({
        history: createWebHistory(),
        routes: routes,
    });
};

const router = buildRouter();

// router.beforeEach((to, from, next) => {
//     const userStore = useUserStore();
//     if (!userStore.hasFetched) {
//         if (userStore.loading === null || userStore.loading === undefined) {
//             userStore.loading = userStore.fetchUser();
//         }
//         userStore.loading.then(() => {
//             check(to, from, next);
//         });
//     } else {
//         check(to, from, next);
//     }
// });

export default router;