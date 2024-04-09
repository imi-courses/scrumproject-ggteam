import { setEngine } from "crypto";
import { FormEvent, useEffect, useState } from "react";
import { Button } from "ui/button";
import { Input } from "ui/input";

const CreateEmployeeForm = () => {
  const [email, setEmail] = useState("");
  const [firstname, setFirstname] = useState("");
  const [surname, setSurname] = useState("");
  const [middlename, setMiddlename] = useState("");
  const [isAuth, setAuth] = useState(false);

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault();
    const response = await fetch(
      import.meta.env.VITE_API_URL + "/employee/create",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + localStorage.getItem("access_token"),
        },
        body: JSON.stringify({
          email,
          firstname,
          surname,
          middlename,
        }),
        credentials: "include",
      },
    );

    const json = await response.json();
    console.log(json);
  };

  useEffect(() => {
    if (window && localStorage.getItem("access_token")) {
      setAuth(true);
    }
  }, []);
  if (!isAuth) return;

  return (
    <form onSubmit={onSubmit}>
      <div>
        <span>ФИО</span>
        <div>
          <Input value={surname} onChange={(e) => setSurname(e.target.value)} />
          <Input
            value={firstname}
            onChange={(e) => setFirstname(e.target.value)}
          />
          <Input
            value={middlename}
            onChange={(e) => setMiddlename(e.target.value)}
          />
        </div>
      </div>

      <Input value={email} onChange={(e) => setEmail(e.target.value)} />
      <Button type="submit">Авторизоваться</Button>
    </form>
  );
};

export default CreateEmployeeForm;
