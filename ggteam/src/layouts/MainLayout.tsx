import { FC, PropsWithChildren } from "react";
import { useAuth } from "providers/auth";
import Header from "@/widgets/Header";
import { Navigate, useLocation } from "react-router-dom";

const MainLayout: FC<PropsWithChildren> = ({ children }) => {
  const { isLoading: authIsLoading, isAuth } = useAuth();
  const location = useLocation();
  if (authIsLoading)
    return (
      <main className="w-screen h-screen flex justify-center items-center">
        LOADING. . .
      </main>
    );
  if (isAuth)
    return (
      <main>
        <Header />
        {children}
      </main>
    );
  if (location.pathname == "/auth") {
    return (
      <main className="w-screen h-screen bg-gradient-to-tl from-slate-400 to-slate-50">
        {children}
      </main>
    );
  }
  return <Navigate to={"/auth"} replace />;
};

export default MainLayout;
