import { FC } from "react";
import { useAuth } from "src/app/providers/auth";
import Button from "ui/Button";

const EmployeeDashboardPage: FC = () => {
  const { logout } = useAuth();
  return (
    <div>
      <Button onClick={logout}>LogOUT</Button>
    </div>
  );
};

export default EmployeeDashboardPage;
