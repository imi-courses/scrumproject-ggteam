import { ColumnDef } from "@tanstack/react-table";

export type Client = {
  id: string;
  fio: string;
  email: string;
  phone: string;
};

export const columns: ColumnDef<Client>[] = [
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
  {
    accessorKey: "phone",
    header: "Номер телефона",
  },
];
