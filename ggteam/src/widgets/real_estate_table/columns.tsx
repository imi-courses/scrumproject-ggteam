import { ColumnDef } from "@tanstack/react-table";

export type RealEstate = {
  id: string;
  address: string;
  type: string;
  client_id: string;
};

export const columns: ColumnDef<RealEstate>[] = [
  {
    accessorKey: "id",
    header: "id",
  },
  {
    accessorKey: "address",
    header: "Адрес",
  },
  {
    accessorKey: "type",
    header: "Тип",
  },
  {
    accessorKey: "client_id",
    header: "ID клиента",
  },
];
