import { useAuth } from "@/app/providers/auth";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/shared/ui/form";
import { useToast } from "@/shared/ui/use-toast";
import { zodResolver } from "@hookform/resolvers/zod";
import { FC } from "react";
import { useForm } from "react-hook-form";
import { Button } from "ui/button";
import { Input } from "ui/input";
import { z } from "zod";
import {
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/shared/ui/dialog";

interface EditEmployeeFormProps {
  employeeData: {
    id: string;
    email: string;
    fio: string;
  };
  getEmployees: () => void;
}

const EditEmployeeForm: FC<EditEmployeeFormProps> = ({
  employeeData,
  getEmployees,
}) => {
  const { token } = useAuth();
  const { toast } = useToast();

  const formSchema = z.object({
    email: z.string().email(),
    firstname: z.string().min(1).max(16),
    surname: z.string().min(1).max(16),
    middlename: z.string().optional(),
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: employeeData.email,
      firstname: employeeData.fio.split(" ")[1],
      surname: employeeData.fio.split(" ")[0],
      middlename:
        employeeData.fio.split(" ").length == 3
          ? employeeData.fio.split(" ")[2]
          : "",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    const response = await fetch(
      import.meta.env.VITE_API_URL + "/employee/" + employeeData.id,
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token,
        },
        body: JSON.stringify(values),
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
          ? "Сотрудник был успешно обновлен"
          : "Что-то пошло не так",
      description: response.status === 200 ? "" : json["message"],
    });
  };

  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Редактирование сотрудника</DialogTitle>
        <DialogDescription>Ошибся? Исправь.</DialogDescription>
      </DialogHeader>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Электронная почта</FormLabel>
                <FormControl>
                  <Input placeholder="Введите почту" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="surname"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Фамилия</FormLabel>
                <FormControl>
                  <Input placeholder="Введите фамилию" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="firstname"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Имя</FormLabel>
                <FormControl>
                  <Input placeholder="Введите имя" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="middlename"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Отчетство</FormLabel>
                <FormControl>
                  <Input placeholder="Введите отчетство" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <DialogFooter className="py-4">
            <Button type="submit">Update</Button>
          </DialogFooter>
        </form>
      </Form>
    </DialogContent>
  );
};

export default EditEmployeeForm;
