import { useAuth } from "@/app/providers/auth";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/shared/ui/card";
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
import * as Dialog from '@radix-ui/react-dialog';

interface EditEmployeeFormProps {
  employeeData: {
    email: string;
    firstname: string;
    surname: string;
    middlename?: string;
  };
  onClose: () => void; // Function to close the dialog
}

const EditEmployeeForm: FC<EditEmployeeFormProps> = ({
  employeeData,
  onClose,
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
      firstname: employeeData.firstname,
      surname: employeeData.surname,
      middlename: employeeData.middlename || "",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    const response = await fetch(
      import.meta.env.VITE_API_URL + "/employee/update",
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
    toast({
      title:
        response.status === 200
          ? "Employee updated successfully"
          : "Something went wrong",
      description: response.status === 200 ? "" : json["message"],
    });
    onClose(); // Close the dialog after updating
    console.log(json);
  };

  return (
    <Dialog.Root onClose={onClose}>
      <Dialog.Trigger />
      <Dialog.Content>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <Card>
              <CardHeader>
                <CardTitle>Edit Employee</CardTitle>
                <CardDescription>
                  Update employee information here
                </CardDescription>
              </CardHeader>
              <CardContent>
                <FormField
                  control={form.control}
                  name="email"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Email</FormLabel>
                      <FormControl>
                        <Input placeholder="Email" {...field} />
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
                      <FormLabel>Surname</FormLabel>
                      <FormControl>
                        <Input placeholder="Surname" {...field} />
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
                      <FormLabel>First Name</FormLabel>
                      <FormControl>
                        <Input placeholder="First Name" {...field} />
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
                      <FormLabel>Middle Name</FormLabel>
                      <FormControl>
                        <Input placeholder="Middle Name" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </CardContent>
              <CardFooter>
                <Button type="submit" className="w-full">
                  Update
                </Button>
              </CardFooter>
            </Card>
          </form>
        </Form>
      </Dialog.Content>
    </Dialog.Root>
  );
};

export default EditEmployeeForm;
