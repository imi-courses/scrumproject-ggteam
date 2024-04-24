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

interface EditRealEstateFormProps {
  realEstateData: {
    id: string;
    address: string;
    type: string;
    clientId: string;
  };
  getRealEstates: () => void;
}

const EditRealEstateForm: FC<EditRealEstateFormProps> = ({
  realEstateData,
  getRealEstates,
}) => {
  const { token } = useAuth();
  const { toast } = useToast();

  const formSchema = z.object({
    address: z.string(),
    type: z.string(),
    clientId: z.string().uuid(),
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      address: realEstateData.address,
      type: realEstateData.type,
      clientId: realEstateData.clientId,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    const response = await fetch(
      import.meta.env.VITE_API_URL + "/real-estate/" + realEstateData.id,
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
      getRealEstates();
    }
    toast({
      variant: response.status === 200 ? "default" : "destructive",
      title:
        response.status === 200
          ? "Недвижимость была успешно обновлена"
          : "Что-то пошло не так",
      description: response.status === 200 ? "" : json["message"],
    });
  };

  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Редактирование недвижимости</DialogTitle>
        <DialogDescription>Ошибся? Исправь.</DialogDescription>
      </DialogHeader>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField
            control={form.control}
            name="address"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Адрес</FormLabel>
                <FormControl>
                  <Input placeholder="Введите адрес" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="type"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Тип</FormLabel>
                <FormControl>
                  <Input placeholder="Введите тип" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="clientId"
            render={({ field }) => (
              <FormItem>
                <FormLabel>ID клиента</FormLabel>
                <FormControl>
                  <Input placeholder="Введите id клиента" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <DialogFooter className="py-4">
            <Button type="submit">Сохранить</Button>
          </DialogFooter>
        </form>
      </Form>
    </DialogContent>
  );
};

export default EditRealEstateForm;
