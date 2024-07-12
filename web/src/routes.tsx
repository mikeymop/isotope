import { RouteObject, createBrowserRouter } from "react-router-dom";
import { Outlet } from "react-router-dom";
import MainLayout from "./layouts/MainLayout";
import { Root, Gallery, Photos, Customers } from "@/pages";

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
        {
          path: "/photos",
          element: <Photos />,
        },
        {
          path: "/customers",
          element: <Customers />,
        },
      ],
    },
  ] as RouteObject[],
  {},
);
