import { useAuth } from "@/app/providers/auth";
import { FC, PropsWithChildren } from "react";
import { Navigate } from "react-router-dom";

const AdminLayout: FC<PropsWithChildren> = ({ children }) => {
  const { userRole } = useAuth();

  if (userRole != "admin") return <Navigate to="/" replace />;
  return children;
};

export default AdminLayout;
