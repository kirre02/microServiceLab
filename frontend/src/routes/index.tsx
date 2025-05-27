import {createRootRoute, createRoute, createRouter} from '@tanstack/react-router';
import Callback from './callback';
import { errorPage } from "./errorPage/errorPage.tsx";


const rootRoute = createRootRoute({});

const callbackRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/callback',
    component: Callback,
});

export const errorRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/error',
    component: errorPage,
});


const routeTree = rootRoute.addChildren([
    callbackRoute,
    errorRoute
]);

export const router = createRouter({
    routeTree
})

// Register your router for maximum type safety
declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router
    }
}

