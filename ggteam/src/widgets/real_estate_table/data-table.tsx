import { useState } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/shared/ui/table";
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/shared/ui/dropdown-menu";
import { Button } from "@/shared/ui/button";
import { Dialog, DialogTrigger } from "@/shared/ui/dialog";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/shared/ui/alert-dialog";
import { useAuth } from "@/app/providers/auth";
import { toast } from "@/shared/ui/use-toast";
import EditRealEstateForm from "@/features/EditRealEstateForm";
import EditRealEstatePDF from "@/features/EditRealEstatePDF";
import SelectRealEstatePDF from "@/features/SelectRealEstatePDF";

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[];
  data: TData[];
  getRealEstates: () => void;
}

export function DataTable<TData, TValue>({
  columns,
  data,
  getRealEstates,
}: DataTableProps<TData, TValue>) {
  const [dialogus, setDialogus] = useState("edit");

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  const { token, getUserInfo } = useAuth();

  const deleteRealEstate = async (id: string) => {
    const response = await fetch(
      import.meta.env.VITE_API_URL + "/real-estate/" + id,
      {
        method: "DELETE",
        headers: {
          Authorization: "Bearer " + token,
        },
        credentials: "include",
      },
    );

    const json = await response.json();
    if (response.status === 200) {
      getRealEstates();
    }
    if (response.status === 401) {
      getUserInfo();
    }
    toast({
      variant: response.status === 200 ? "default" : "destructive",
      title:
        response.status === 200
          ? "Недвижимость была успешно удалена"
          : "Что-то пошло не так",
      description: response.status === 200 ? "" : json["message"],
    });
  };

  return (
    <div>
      <Table>
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => {
                return (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                        header.column.columnDef.header,
                        header.getContext(),
                      )}
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row) => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && "selected"}
              >
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
                <TableCell>
                  <AlertDialog>
                    <Dialog>
                      <DropdownMenu>
                        <DropdownMenuTrigger>
                          <Button variant="ghost" size="sm">
                            ⋮
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent>
                          <DropdownMenuItem>
                            <DialogTrigger
                              onClick={() => setDialogus("select pdf")}
                            >
                              Справка о недвижимости
                            </DialogTrigger>
                          </DropdownMenuItem>
                          <DropdownMenuItem>
                            <DialogTrigger
                              onClick={() => setDialogus("update pdf")}
                            >
                              Форма заявления об изменении характеристик
                              недвижимости
                            </DialogTrigger>
                          </DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem>
                            <DialogTrigger onClick={() => setDialogus("edit")}>
                              Редактировать
                            </DialogTrigger>
                          </DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem>
                            <AlertDialogTrigger>Удалить</AlertDialogTrigger>
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                      {dialogus == "select pdf" && (
                        <SelectRealEstatePDF
                          address={row.getValue("address")}
                          type={row.getValue("type")}
                          clientId={row.getValue("client_id")}
                        />
                      )}

                      {dialogus == "update pdf" && (
                        <EditRealEstatePDF
                          address={row.getValue("address")}
                          type={row.getValue("type")}
                          clientId={row.getValue("client_id")}
                        />
                      )}

                      {dialogus == "edit" && (
                        <EditRealEstateForm
                          realEstateData={{
                            id: row.getValue("id"),
                            address: row.getValue("address"),
                            type: row.getValue("type"),
                            clientId: row.getValue("client_id"),
                          }}
                          getRealEstates={getRealEstates}
                        />
                      )}
                      <AlertDialogContent>
                        <AlertDialogHeader>
                          <AlertDialogTitle>Вы уверены?</AlertDialogTitle>
                          <AlertDialogDescription>
                            Вы навсегда потеряете данные об этом сотруднике
                          </AlertDialogDescription>
                        </AlertDialogHeader>
                        <AlertDialogFooter>
                          <AlertDialogCancel>Отмена</AlertDialogCancel>
                          <AlertDialogAction
                            onClick={() => deleteRealEstate(row.getValue("id"))}
                          >
                            Удалить
                          </AlertDialogAction>
                        </AlertDialogFooter>
                      </AlertDialogContent>
                    </Dialog>
                  </AlertDialog>
                </TableCell>
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
