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
import EditEmployeeForm from "@/features/EditEmployeeForm";
import { Dialog } from "@/shared/ui/dialog";
import { DialogTrigger } from "@radix-ui/react-dialog";
import { Button } from "@/shared/ui/button";
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

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[];
  data: TData[];
  getEmployees: () => void;
}

export function DataTable<TData, TValue>({
  columns,
  data,
  getEmployees,
}: DataTableProps<TData, TValue>) {
  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  const { token } = useAuth();

  const deleteEmployee = async (id: string) => {
    const response = await fetch(
      import.meta.env.VITE_API_URL + "/employee/" + id,
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
      getEmployees();
    }
    toast({
      variant: response.status === 200 ? "default" : "destructive",
      title:
        response.status === 200
          ? "Сотрудник был успешно удален"
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
                            <DialogTrigger>Редактировать</DialogTrigger>
                          </DropdownMenuItem>{" "}
                          <DropdownMenuSeparator />
                          <DropdownMenuItem>
                            <AlertDialogTrigger>Удалить</AlertDialogTrigger>
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>

                      <EditEmployeeForm
                        employeeData={{
                          id: row.getValue("id"),
                          email: row.getValue("email"),
                          fio: row.getValue("fio"),
                        }}
                        getEmployees={getEmployees}
                      />
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
                            onClick={() => deleteEmployee(row.getValue("id"))}
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
