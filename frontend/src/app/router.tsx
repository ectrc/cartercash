import {
  createRootRoute,
  createRoute,
  createRouter,
  Navigate,
} from "@tanstack/react-router";

import Frame from "./frame";
import Browser from "./browser";
import Game from "./game";

export const rootRoute = createRootRoute({
  component: Frame,
  notFoundComponent: () => <Navigate to="" />,
});

export const browserRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: Browser,
});

export const gameRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/game/$id",
  component: Game,
});

const router = createRouter({
  routeTree: rootRoute.addChildren([browserRoute, gameRoute]),
});

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

export default router;
