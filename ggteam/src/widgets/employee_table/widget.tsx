import { FC } from "react";
import { Employee, columns } from "./columns";
import { DataTable } from "./data-table";

interface EmployeeTableProps {
  data: Employee[];
  getEmployees: () => void;
}

const EmployeeTable: FC<EmployeeTableProps> = ({ data, getEmployees }) => {
  return (
    <DataTable columns={columns} data={data} getEmployees={getEmployees} />
  );
};

export default EmployeeTable;
