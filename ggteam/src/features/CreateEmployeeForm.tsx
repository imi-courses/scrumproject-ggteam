import { FormEvent, useEffect, useState } from "react";
import Button from "ui/Button";
import Input from "ui/Input";

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
          <Input value={surname} setValue={setSurname} />
          <Input value={firstname} setValue={setFirstname} />
          <Input value={middlename} setValue={setMiddlename} />
        </div>
      </div>

      <Input value={email} setValue={setEmail} />
      <Button type="submit">Авторизоваться</Button>
    </form>
  );
};

export default CreateEmployeeForm;
