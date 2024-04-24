import { useAuth } from "@/app/providers/auth";
import CreateClient from "@/features/CreateClientForm";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/shared/ui/tabs";
import { Client } from "@/widgets/client_table/columns";
import ClientTable from "@/widgets/client_table/widget";
import { RealEstate } from "@/widgets/real_estate_table/columns";
import RealEstateTable from "@/widgets/real_estate_table/widget";
import { FC, useCallback, useEffect, useState } from "react";

const EmployeeDashboardPage: FC = () => {
  const [clientData, setClientData] = useState<Client[]>([]);
  const [realEstateData, setrealEstateData] = useState<RealEstate[]>([]);
  const { token } = useAuth();

  const getClientData = useCallback(async () => {
    return fetch(
      import.meta.env.VITE_API_URL + "/client/find-all?page=1&count=10",
      {
        headers: {
          Authorization: "Bearer " + token,
        },
      },
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
        setClientData(clients);
      }
    });
  }, [token]);

  const getRealEstateData = useCallback(async () => {
    return fetch(
      import.meta.env.VITE_API_URL + "/real-estate/find-all?page=1&count=10",
      {
        headers: {
          Authorization: "Bearer " + token,
        },
      },
    ).then(async (res) => {
      if (res.status === 200) {
        const json = await res.json();
        const realEstates: RealEstate[] = [];
        for (let i = 0; i < json["real_estates"].length; i++) {
          const realEstate = json["real_estates"][i];
          realEstates.push({
            id: realEstate.id,
            address: realEstate.address,
            type: realEstate.type,
            client_id: realEstate.client_id,
          });
        }
        setrealEstateData(realEstates);
      }
    });
  }, [token]);

  useEffect(() => {
    getClientData();
    getRealEstateData();
  }, [getClientData, getRealEstateData]);

  return (
    <section className="container sm:flex">
      <Tabs
        defaultValue="clients"
        className="sm:flex w-full sm:gap-4 sm:justify-start sm:items-start"
      >
        <TabsList className="flex-col h-20 w-full sm:w-auto">
          <TabsTrigger value="clients" className="w-full">
            Список клиентов
          </TabsTrigger>
          <TabsTrigger value="real-estates" className="w-full">
            Список недвижимости
          </TabsTrigger>
        </TabsList>
        <TabsContent value="clients" className="sm:mt-0 w-full">
          <ClientTable
            data={clientData}
            getClients={getClientData}
            getRealEstates={getRealEstateData}
          />
          <CreateClient getClients={getClientData} />
        </TabsContent>
        <TabsContent value="real-estates" className="sm:mt-0 w-full">
          <RealEstateTable
            data={realEstateData}
            getRealEstates={getRealEstateData}
          />
        </TabsContent>
      </Tabs>
    </section>
  );
};

export default EmployeeDashboardPage;
