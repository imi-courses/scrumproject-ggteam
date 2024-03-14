import React from "react";
import ReactDOM from "react-dom/client";
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import "styles/app.css";
import Root from "src/routes/root";
import Admin from "src/routes/admin";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
  },
  {
    path: "admin",
    element: <Admin />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
);
