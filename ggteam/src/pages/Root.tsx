import { FC } from "react";
import { Outlet } from "react-router-dom";
import { AuthProvider } from "providers/auth";
import MainLayout from "@/layouts/MainLayout";

const Root: FC = () => {
  return (
    <AuthProvider>
      <MainLayout>
        <Outlet />
      </MainLayout>
    </AuthProvider>
  );
};

export default Root;
