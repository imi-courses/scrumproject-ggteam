import { FC } from "react";
import "styles/Header.css";

const Header: FC = () => {
  return (
    <header className="headerStyle">
      <div className="headerLogo">
        <h3>Росреестр</h3>
      </div>
      <div className="headerProfile">
        <p>Войти</p>
        {/* <p className="headerProfile">Профиль</p> */}
      </div>
    </header>
  );
};

export default Header;
