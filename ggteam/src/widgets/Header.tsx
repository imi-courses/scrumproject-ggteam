import { useAuth } from "@/app/providers/auth";
import { Button } from "@/shared/ui/button";
import { FC } from "react";
import "styles/Header.css";

const Header: FC = () => {
  const { logout } = useAuth();
  return (
    <header className="bg-gradient-to-l from-slate-300 to-slate-50 mb-4">
      <div className="container py-2 flex justify-between items-center">
        <a className="text-2xl font-semibold uppercase" href="/">
          Росреестр
        </a>
        <Button onClick={logout}>Выйти</Button>
      </div>
    </header>
  );
};

export default Header;
