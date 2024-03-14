import { AuthProvider } from "src/app/providers/auth";
import AuthPage from "src/pages/Auth";

const MainLayout = () => {
  return (
    <>
      <AuthProvider>
        <AuthPage />
      </AuthProvider>
    </>
  );
};

export default MainLayout;
