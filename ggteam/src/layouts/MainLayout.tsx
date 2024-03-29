import { FC, PropsWithChildren } from "react";
import { useAuth } from "src/app/providers/auth";

const MainLayout: FC<PropsWithChildren> = ({ children }) => {
  const { isLoading: authIsLoading } = useAuth();
  return <>{authIsLoading ? <div>Loading</div> : children}</>;
};

export default MainLayout;
