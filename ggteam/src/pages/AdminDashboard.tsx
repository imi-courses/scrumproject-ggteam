import { FC, useCallback, useEffect, useState } from "react";
import CreateEmployeeForm from "@/features/CreateEmployeeForm";
import { useAuth } from "@/app/providers/auth";
import { Navigate } from "react-router-dom";
import { Tabs, TabsList, TabsContent, TabsTrigger } from "@/shared/ui/tabs";
import EmployeeTable from "@/widgets/employee_table/widget";
import { Employee } from "@/widgets/employee_table/columns";

const AdminDashboardPage: FC = () => {
  const [data, setData] = useState<Employee[]>([]);
  const { userRole, token } = useAuth();

  const getData = useCallback(async () => {
    return fetch(
      import.meta.env.VITE_API_URL + "/employee/find-all?page=1&count=10",
      {
        headers: {
          Authorization: "Bearer " + token,
        },
      },
    ).then(async (res) => {
      if (res.status === 200) {
        const json = await res.json();
        const employees: Employee[] = [];
        for (let i = 0; i < json["employees"].length; i++) {
          const employee = json["employees"][i];
          employees.push({
            id: employee.id,
            fio: `${employee.surname} ${employee.firstname} ${employee.middlename}`,
            email: employee.email,
          });
        }
        setData(employees);
      }
    });
  }, [token]);

  useEffect(() => {
    getData();
  }, [getData]);

  if (userRole != "admin") return <Navigate to="/" replace />;
  return (
    <section className="">
      <Tabs
        defaultValue="list"
        className="sm:flex flex-col w-full sm:gap-4 sm:justify-start sm:items-start"
      >
        <TabsList className="h-10 w-full bg-[#035835] flex justify-between">
          <TabsTrigger value="list" className="w-full">
              Список сотрудников
          </TabsTrigger>
          <TabsTrigger value="create" className="w-full">
              Регистрация сотрудника
          </TabsTrigger>
        </TabsList>
        <section className="container sm:flex">
          <TabsContent value="list" className="flex-grow bg-[#E8E8E8]" style={{
                    borderTopLeftRadius: "0.5rem",
                    borderTopRightRadius: "0.5rem",
                  }}>
            <EmployeeTable data={data} getEmployees={getData} />
          </TabsContent>
          <TabsContent value="create" className="flex-grow">
            <CreateEmployeeForm getEmployees={getData} />
          </TabsContent>
        </section>     
      </Tabs>
    </section>
  );
};

export default AdminDashboardPage;
