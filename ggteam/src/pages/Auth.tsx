import { useState } from "react";
import { useAuth } from "src/app/providers/auth";
import AdminForm from "src/widgets/AdminForm";
import CreateEmployeeForm from "src/widgets/CreateEmployeeForm";
import EmployeeForm from "src/widgets/EmployeeForm";
import Button from "ui/Button";

interface showAuthForm {
  AdminAuthForm: boolean;
  EmployeeAuthForm: boolean;
}

const AuthPage = () => {
  const [showAuthForm, setAuthForm] = useState<showAuthForm>({
    AdminAuthForm: false,
    EmployeeAuthForm: true,
  });

  const { isAuth } = useAuth();

  const changeAuthForm = (adminForm: boolean) => {
    if (adminForm) {
      setAuthForm({ AdminAuthForm: true, EmployeeAuthForm: false });
    } else {
      setAuthForm({ AdminAuthForm: false, EmployeeAuthForm: true });
    }
  };

  const test = async () => {
    const response = await fetch(import.meta.env.VITE_API_URL + "/auth/", {
      credentials: "include",
    });

    const json = await response.json();
    console.log(json);
  };

  if (isAuth) return;

  return (
    <section>
      <h1 onClick={test}>Авторизация</h1>
      <span>Росреестр</span>
      <div>
        <div>
          <Button onClick={() => changeAuthForm(false)}>Работник</Button>
          <Button onClick={() => changeAuthForm(true)}>Администратор</Button>
        </div>
        {showAuthForm.AdminAuthForm && <AdminForm />}
        {showAuthForm.EmployeeAuthForm && <EmployeeForm />}
        <CreateEmployeeForm />
      </div>
    </section>
  );
};

export default AuthPage;
