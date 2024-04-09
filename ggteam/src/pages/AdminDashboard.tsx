import { FC } from "react";
import SearchEmployeeForm from "@/features/SearchEmployeeForm";
import CreateEmployeeForm from "@/features/CreateEmployeeForm";

const AdminDashboardPage: FC = () => {
  return (
    <section>
      <CreateEmployeeForm />
      <SearchEmployeeForm />
    </section>
  );
};

export default AdminDashboardPage;
