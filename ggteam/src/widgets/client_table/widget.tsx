import { FC } from "react";
import { Client, columns } from "./columns";
import { DataTable } from "./data-table";

interface ClientTableProps {
  data: Client[];
  getClients: () => void;
  getRealEstates: () => void;
}

const ClientTable: FC<ClientTableProps> = ({
  data,
  getClients,
  getRealEstates,
}) => {
  return (
    <DataTable
      columns={columns}
      data={data}
      getClients={getClients}
      getRealEstates={getRealEstates}
    />
  );
};

export default ClientTable;
