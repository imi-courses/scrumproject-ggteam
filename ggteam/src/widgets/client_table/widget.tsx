import { FC } from "react";
import { Client, columns } from "./columns";
import { DataTable } from "./data-table";

interface ClientTableProps {
  data: Client[];
  getClients: () => void;
}

const ClientTable: FC<ClientTableProps> = ({ data, getClients }) => {
  return (
    <DataTable columns={columns} data={data} getClients={getClients} />
  );
};

export default ClientTable;
