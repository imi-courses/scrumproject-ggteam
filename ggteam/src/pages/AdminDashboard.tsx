import { FC } from "react";
import SearchEmployeeForm from "@/features/SearchEmployeeForm";
import CreateEmployeeForm from "@/features/CreateEmployeeForm";
import { useAuth } from "@/app/providers/auth";
import { Navigate } from "react-router-dom";

const AdminDashboardPage: FC = () => {
  const { userRole } = useAuth();

  if (userRole != "admin") return <Navigate to="/" replace />;
  return (
    <section>
      <CreateEmployeeForm />
      <SearchEmployeeForm />
    </section>
  );
};

export default AdminDashboardPage;
