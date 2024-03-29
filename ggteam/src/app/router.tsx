import { FC } from "react";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import AdminDashboardPage from "src/pages/AdminDashboard";
import AuthPage from "src/pages/Auth";
import EmployeeDashboardPage from "src/pages/EmployeeDashboard";
import Root from "src/pages/Root";

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
