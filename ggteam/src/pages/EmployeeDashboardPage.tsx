import { FC, useCallback, useEffect, useState } from "react";
import CreateClientForm from "@/features/CreateClientForm";
import { useAuth } from "@/app/providers/auth";
import { Tabs, TabsList, TabsContent, TabsTrigger } from "@/shared/ui/tabs";
import ClientTable from "@/widgets/client_table/widget";
import { Client } from "@/widgets/client_table/columns";

export const EmployeeDashboardPage: FC = () => {
  const [data, setData] = useState<Client[]>([]);
  const { userRole, token } = useAuth();

  const getData = useCallback(async () => {
    return fetch(
      import.meta.env.VITE_API_URL + "/client/?page=1&count=10",
      {
        method: "GET",
        headers: {
          Authorization: "Bearer " + token,
        },
        credentials: "include"
      }
    ).then(async (res) => {
      if (res.status === 200) {
        const json = await res.json();
        const clients: Client[] = [];
        for (let i = 0; i < json["clients"].length; i++) {
          const client = json["clients"][i];
          clients.push({
            id: client.id,
            fio: `${client.surname} ${client.firstname} ${client.middlename}`,
            email: client.email,
            phone: client.phone,
          });
        }
        setData(clients);
      }
    });
  }, [token]);

  useEffect(() => {
    getData();
  }, [getData]);

  return (
    <section className="container sm:flex">
      <Tabs
        defaultValue="list"
        className="sm:flex w-full sm:gap-4 sm:justify-start sm:items-start"
      >
        <TabsList className="flex-col h-20 w-full sm:w-auto">
          <TabsTrigger value="list" className="w-full">
            Список клиентов
          </TabsTrigger>
          <TabsTrigger value="create" className="w-full">
            Список
          </TabsTrigger>
        </TabsList>
        <TabsContent value="list" className="sm:mt-0 w-full">
          <ClientTable data={data} getClients={getData} />
        </TabsContent>
        <TabsContent value="create" className="sm:mt-0 w-full">
          <CreateClientForm getClients={getData} />
        </TabsContent>
      </Tabs>
    </section>
  );
};
