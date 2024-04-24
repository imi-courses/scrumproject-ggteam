import { useAuth } from "@/app/providers/auth";
import { Button } from "@/shared/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/shared/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/shared/ui/form";
import { Input } from "@/shared/ui/input";
import { useToast } from "@/shared/ui/use-toast";
import { zodResolver } from "@hookform/resolvers/zod";
import { FC } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

interface CreateRealEstateProps {
  getRealEstates: () => void;
  clientId: string;
}

const formSchema = z.object({
  address: z.string(),
  type: z.string(),
  client_id: z.string().uuid(),
});

const CreateRealEstate: FC<CreateRealEstateProps> = ({
  getRealEstates,
  clientId,
}) => {
  const { token } = useAuth();
  const { toast } = useToast();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      address: "",
      type: "",
      client_id: clientId,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    const response = await fetch(
      import.meta.env.VITE_API_URL + "/real-estate/create",
      {
        method: "POST",
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
          ? "Недвижимость была успешно добавлена"
          : "Что-то пошло не так",
      description: response.status === 200 ? "" : json["message"],
    });
  };
  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Добавление недвижимости</DialogTitle>
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
                <FormLabel>Тип недвижимости</FormLabel>
                <FormControl>
                  <Input placeholder="Введите тип" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <DialogFooter className="py-4">
            <Button type="submit">Добавить</Button>
          </DialogFooter>
        </form>
      </Form>
    </DialogContent>
  );
};

export default CreateRealEstate;
