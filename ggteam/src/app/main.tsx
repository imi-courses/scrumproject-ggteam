import React from "react";
import ReactDOM from "react-dom/client";
import MainLayout from "src/layouts/MainLayout";
import "styles/app.css";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <MainLayout />
  </React.StrictMode>,
);
