import { FC } from "react";
import { RealEstate, columns } from "./columns";
import { DataTable } from "./data-table";

interface RealEstateTableProps {
  data: RealEstate[];
  getRealEstates: () => void;
}

const RealEstateTable: FC<RealEstateTableProps> = ({
  data,
  getRealEstates,
}) => {
  return (
    <DataTable columns={columns} data={data} getRealEstates={getRealEstates} />
  );
};

export default RealEstateTable;
