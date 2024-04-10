import { ColumnDef } from "@tanstack/react-table";

export type Employee = {
  id: string;
  fio: string;
  email: string;
};

export const columns: ColumnDef<Employee>[] = [
  {
    accessorKey: "id",
    header: "id",
  },
  {
    accessorKey: "fio",
    header: "ФИО",
  },
  {
    accessorKey: "email",
    header: "Электронная почта",
  },
];
