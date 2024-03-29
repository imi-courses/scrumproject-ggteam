import { FormEvent, useState } from "react";
import { useAuth } from "src/app/providers/auth";
import Button from "ui/Button";
import Input from "ui/Input";
import "styles/Auth.css";

const AdminForm = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { setAuth } = useAuth();

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault();
    const response = await fetch(import.meta.env.VITE_API_URL + "/auth/admin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email: email, password: password }),
      credentials: "include",
    });
    const json = await response.json();
    if (response.status === 200) {
      localStorage.setItem("access_token", json["access_token"]);
      setAuth(true);
    }
    console.log(json);
  };

  return (
    <form onSubmit={onSubmit} className="classAdminForm">
      <h3 className="classMiniTitleAuth">Авторизация Администратора</h3>
      <div className="classFormGroup">
        <label htmlFor="email" className="classFormLabel">
          Email
        </label>
        <Input value={email} setValue={setEmail} />
      </div>
      <div className="classFormGroup">
        <label htmlFor="password" className="classFormLabel">
          Password
        </label>
        <Input value={password} setValue={setPassword} />
      </div>
      <Button type="submit">Авторизоваться</Button>
    </form>
  );
};

export default AdminForm;
