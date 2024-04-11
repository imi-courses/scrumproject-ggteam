import { useAuth } from "providers/auth";
import { Button } from "ui/button";
import { Input } from "ui/input";
import "styles/Auth.css";

import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/shared/ui/form";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/shared/ui/card";
import { useNavigate } from "react-router-dom";
import { toast } from "@/shared/ui/use-toast";

const formSchema = z.object({
  email: z.string().email(),
  password: z.string(),
});

const AdminForm = () => {
  const { setAuth, setToken, setUserRole } = useAuth();
  const navigate = useNavigate();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      password: "",
      email: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    const response = await fetch(import.meta.env.VITE_API_URL + "/auth/admin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(values),
      credentials: "include",
    });
    const json = await response.json();
    if (response.status === 200) {
      localStorage.setItem("access_token", json["access_token"]);
      setAuth(true);
      setToken(json["access_token"]);
      setUserRole("admin");
      navigate("/admin");
    } else {
      toast({
        variant: "destructive",
        title: "Что-то пошло не так",
        description: json.message,
      });
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <Card>
          <CardHeader>
            <CardTitle>Авторизация Администратора</CardTitle>
            <CardDescription>
              Только истинный админ сможет войти
            </CardDescription>
          </CardHeader>
          <CardContent>
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Электронная почта</FormLabel>
                  <FormControl>
                    <Input placeholder="Почта" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Секретный ключ</FormLabel>
                  <FormControl>
                    <Input placeholder="Ключ" type="password" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </CardContent>
          <CardFooter>
            <Button className="w-full" type="submit">
              Авторизоваться
            </Button>
          </CardFooter>
        </Card>
      </form>
    </Form>
  );
};

export default AdminForm;
