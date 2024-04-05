import { FC } from "react";
import SearchEmployeeForm from "src/features/SearchEmployeeForm";
import CreateEmployeeForm from "src/features/CreateEmployeeForm";

const AdminDashboardPage: FC = () => {
  return <section>
      <CreateEmployeeForm/>
      <SearchEmployeeForm/> 
    </section>;
};

export default AdminDashboardPage;
