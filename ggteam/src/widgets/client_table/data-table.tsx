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
import EditClientForm from "@/features/EditClientForm";
import CreateRealEstatePDF from "@/features/CreateRealEstatePDF";
import CreateRealEstate from "@/features/CreateRealEstate";
import EditRealEstatePDF from "@/features/EditRealEstatePDF";

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[];
  data: TData[];
  getClients: () => void;
  getRealEstates: () => void;
}

export function DataTable<TData, TValue>({
  columns,
  data,
  getClients,
  getRealEstates,
}: DataTableProps<TData, TValue>) {
  const [dialogus, setDialogus] = useState("edit");

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  const { token, getUserInfo } = useAuth();

  const deleteClient = async (id: string) => {
    const response = await fetch(
      import.meta.env.VITE_API_URL + "/client/" + id,
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
      getClients();
    }
    if (response.status === 401) {
      getUserInfo();
    }
    toast({
      variant: response.status === 200 ? "default" : "destructive",
      title:
        response.status === 200
          ? "Клиент был успешно удален"
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
                              onClick={() => setDialogus("create pdf")}
                            >
                              Форма заявления регистрации недвижимости
                            </DialogTrigger>
                          </DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem>
                            <DialogTrigger
                              onClick={() => setDialogus("create")}
                            >
                              Добавить недвижимость
                            </DialogTrigger>
                          </DropdownMenuItem>
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
                      {dialogus == "create pdf" && (
                        <CreateRealEstatePDF
                          phone={row.getValue("phone")}
                          email={row.getValue("email")}
                          fio={row.getValue("fio")}
                        />
                      )}

                      {dialogus == "update pdf" && (
                        <EditRealEstatePDF
                          phone={row.getValue("phone")}
                          email={row.getValue("email")}
                          fio={row.getValue("fio")}
                        />
                      )}

                      {dialogus == "create" && (
                        <CreateRealEstate
                          clientId={row.getValue("id")}
                          getRealEstates={getRealEstates}
                        />
                      )}
                      {dialogus == "edit" && (
                        <EditClientForm
                          clientData={{
                            phone: row.getValue("phone"),
                            id: row.getValue("id"),
                            email: row.getValue("email"),
                            fio: row.getValue("fio"),
                          }}
                          getClients={getClients}
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
                            onClick={() => deleteClient(row.getValue("id"))}
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
