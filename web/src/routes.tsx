import { RouteObject, createBrowserRouter } from "react-router-dom";
import { Outlet } from "react-router-dom";
import MainLayout from "./layouts/MainLayout";
import { Root } from "./routes/root";
import { Gallery } from "./routes/gallery";

export const router = createBrowserRouter(
  [
    {
      element: (
        <MainLayout>
          <Outlet />
        </MainLayout>
      ),
      children: [
        {
          path: "/",
          element: <Root />,
        },
        {
          path: "/gallery",
          element: <Gallery />,
        },
      ],
    },
  ] as RouteObject[],
  {},
);
