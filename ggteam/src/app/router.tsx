import { FC } from "react";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import AdminDashboardPage from "@/pages/AdminDashboard";
import AuthPage from "@/pages/Auth";
import EmployeeDashboardPage from "@/pages/EmployeeDashboard";
import Root from "@/pages/Root";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        path: "",
        element: <EmployeeDashboardPage />,
      },
      {
        path: "admin",
        element: <AdminDashboardPage />,
      },
      {
        path: "auth",
        element: <AuthPage />,
      },
    ],
  },
]);

const Router: FC = () => <RouterProvider router={router} />;

export default Router;
