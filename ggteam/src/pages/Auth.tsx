import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "providers/auth";
import AdminForm from "@/features/AdminForm";
import EmployeeForm from "@/features/EmployeeForm";
import { Button } from "ui/button";
import "styles/Auth.css";

interface showAuthForm {
  AdminAuthForm: boolean;
  EmployeeAuthForm: boolean;
}

const AuthPage = () => {
  const [showAuthForm, setAuthForm] = useState<showAuthForm>({
    AdminAuthForm: false,
    EmployeeAuthForm: true,
  });
  const { isLoading, isAuth } = useAuth();
  const navigate = useNavigate();

  const changeAuthForm = (adminForm: boolean) => {
    if (adminForm) {
      setAuthForm({ AdminAuthForm: true, EmployeeAuthForm: false });
    } else {
      setAuthForm({ AdminAuthForm: false, EmployeeAuthForm: true });
    }
  };

  useEffect(() => {
    if (!isLoading && isAuth) navigate("/");
  }, [isAuth, isLoading, navigate]);

  return (
    <section className="classAuth">
      <h3 className="classTitleAuth">Авторизация</h3>
      <div>
        <Button onClick={() => changeAuthForm(false)}>Работник</Button>
        <Button onClick={() => changeAuthForm(true)}>Администратор</Button>
      </div>
      {showAuthForm.AdminAuthForm && <AdminForm />}
      {showAuthForm.EmployeeAuthForm && <EmployeeForm />}
    </section>
  );
};

export default AuthPage;
