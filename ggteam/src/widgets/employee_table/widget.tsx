import { FC } from "react";
import { Employee, columns } from "./columns";
import { DataTable } from "./data-table";

interface EmployeeTableProps {
  data: Employee[];
}

const EmployeeTable: FC<EmployeeTableProps> = ({ data }) => {
  return <DataTable columns={columns} data={data} />;
};

export default EmployeeTable;
