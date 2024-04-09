import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "providers/auth";
import AdminForm from "@/features/AdminForm";
import EmployeeForm from "@/features/EmployeeForm";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/shared/ui/tabs";

const AuthPage = () => {
  const { isLoading, isAuth } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (!isLoading && isAuth) navigate("/");
  }, [isAuth, isLoading, navigate]);

  return (
    <section className="flex flex-col justify-center items-center pt-10 px-4">
      <Tabs defaultValue="employee" className="w-full xs:w-[400px]">
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger value="employee">Сотрудник</TabsTrigger>
          <TabsTrigger value="admin">Администратор</TabsTrigger>
        </TabsList>
        <TabsContent value="employee">
          <EmployeeForm />
        </TabsContent>
        <TabsContent value="admin">
          <AdminForm />
        </TabsContent>
      </Tabs>
    </section>
  );
};

export default AuthPage;
