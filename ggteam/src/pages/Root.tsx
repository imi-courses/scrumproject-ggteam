import { FC } from "react";
import { Outlet } from "react-router-dom";
import { AuthProvider } from "src/app/providers/auth";
import MainLayout from "src/layouts/MainLayout";

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
