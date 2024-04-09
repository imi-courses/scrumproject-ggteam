import { FC, PropsWithChildren } from "react";
import { useAuth } from "src/app/providers/auth";
import Header from "src/widgets/Header";

const MainLayout: FC<PropsWithChildren> = ({ children }) => {
  const { isLoading: authIsLoading } = useAuth();
  return (
    <>
      {authIsLoading ? (
        <div>Loading</div>
      ) : (
        <main>
          <Header></Header>
          {children}
        </main>
      )}
    </>
  );
};

export default MainLayout;