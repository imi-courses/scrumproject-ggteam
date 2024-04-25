import { useAuth } from "@/app/providers/auth";
import { Button } from "@/shared/ui/button";
import { FC } from "react";
import "styles/Header.css";

const Header: FC = () => {
 const { logout } = useAuth();
 return (
    <header className="bg-white">
      <div className="py-2 flex justify-between items-center pl-8 pr-8"> 
        <a style={{ color: '#035835' }} className="text-2xl font-semibold uppercase" href="/"> 
          Росреестр
        </a>
        <Button onClick={logout}>Выйти</Button>
      </div>
    </header>
 );
};

export default Header;