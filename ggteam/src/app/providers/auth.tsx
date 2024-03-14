import {
  Dispatch,
  FC,
  PropsWithChildren,
  SetStateAction,
  createContext,
  useCallback,
  useContext,
  useEffect,
  useState,
} from "react";
import { useNavigate } from "react-router-dom";

type UserRole = "admin" | "employee" | undefined;

interface AuthContext {
  isAuth: boolean;
  setAuth: Dispatch<SetStateAction<boolean>>;
  isLoading: boolean;
  userRole: UserRole;
}

const defaultValues: AuthContext = {
  isAuth: false,
  setAuth: () => null,
  isLoading: true,
  userRole: undefined,
};

const Context = createContext(defaultValues);

interface AuthProviderProps extends PropsWithChildren { }

export const AuthProvider: FC<AuthProviderProps> = (props) => {
  const { children } = props;
  const [isAuth, setAuth] = useState<boolean>(defaultValues.isAuth);
  const [isLoading, setLoading] = useState<boolean>(defaultValues.isLoading);
  const [userRole, setUserRole] = useState<UserRole>(defaultValues.userRole);
  const navigate = useNavigate();

  const refreshTokens = useCallback(async () => {
    return fetch(import.meta.env.VITE_API_URL + "/auth/refresh", {
      method: "GET",
      credentials: "include",
    })
      .then(async (res) => {
        const data = await res.json();
        if (res.status === 200) {
          localStorage.setItem("access_token", data["access_token"]);
          return true;
        }
        return false;
      })
      .catch((err) => {
        console.log(err);
        return false;
      });
  }, []);

  const getUserInfo = useCallback(
    async (token: string) => {
      return fetch(import.meta.env.VITE_API_URL + "/auth/me", {
        method: "GET",
        headers: {
          Authorization: "Bearer " + token,
        },
      })
        .then(async (res) => {
          const data = await res.json();
          if (res.status === 200) {
            setAuth(true);
            setUserRole(data["role"]);
          } else {
            const isLoggedOut = await refreshTokens();
            if (isLoggedOut) {
              setAuth(true);
              setUserRole(data["role"]);
            } else {
              setLoading(false);
              setAuth(false);
              navigate("/auth");
            }
          }
          setLoading(false);
          return data;
        })
        .catch((err) => console.error(err));
    },
    [navigate, refreshTokens],
  );

  useEffect(() => {
    const token = localStorage.getItem("access_token");
    return () => {
      if (token) {
        getUserInfo(token);
      } else {
        setLoading(false);
      }
    };
  }, [getUserInfo]);

  const exposed = {
    isAuth,
    setAuth,
    isLoading,
    userRole,
  };

  return <Context.Provider value={exposed}>{children}</Context.Provider>;
};

export const useAuth = () => useContext(Context);
