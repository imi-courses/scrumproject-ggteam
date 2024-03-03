import { FormEvent, useState } from "react";
import Button from "ui/Button";
import Input from "ui/Input";

const AdminForm = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault();
    const response = await fetch(import.meta.env.VITE_API_URL + "/auth/admin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email: email, password: password }),
    });

    const json = await response.json();
    console.log(json);
  };

  return (
    <form onSubmit={onSubmit}>
      <Input value={email} setValue={setEmail} />
      <Input value={password} setValue={setPassword} />
      <Button type="submit">Авторизоваться</Button>
    </form>
  );
};

export default AdminForm;
