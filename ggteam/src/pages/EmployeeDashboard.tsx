import { FC } from "react";
import { useAuth } from "providers/auth";
import { Button } from "ui/button";

const EmployeeDashboardPage: FC = () => {
  const { logout } = useAuth();
  return (
    <div>
      <Button onClick={logout}>LogOUT</Button>
    </div>
  );
};

export default EmployeeDashboardPage;
