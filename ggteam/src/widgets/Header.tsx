import { useAuth } from "@/app/providers/auth";
import { Button } from "@/shared/ui/button";
import { FC } from "react";
import "styles/Header.css";

const Header: FC = () => {
 const { logout } = useAuth();
 return (
    <header>
      <div className="container py-2 flex justify-between items-center"> 
        <a style={{ color: '#035835' }} className="text-2xl font-semibold uppercase" href="/"> 
          Росреестр
        </a>
        <Button onClick={logout}>Выйти</Button>
      </div>
    </header>
 );
};

export default Header;