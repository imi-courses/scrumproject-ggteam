import { FC } from "react";
import { Outlet } from "react-router-dom";
import { AuthProvider } from "providers/auth";
import MainLayout from "@/layouts/MainLayout";
import { Toaster } from "@/shared/ui/toaster";

const Root: FC = () => {
  return (
    <AuthProvider>
      <MainLayout>
        <Outlet />
        <Toaster />
      </MainLayout>
    </AuthProvider>
  );
};

export default Root;
