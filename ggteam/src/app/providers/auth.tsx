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
  logout: () => void;
  token: string;
}

const defaultValues: AuthContext = {
  isAuth: false,
  setAuth: () => null,
  isLoading: true,
  userRole: undefined,
  logout: () => null,
  token: "",
};

const Context = createContext(defaultValues);

interface AuthProviderProps extends PropsWithChildren { }

export const AuthProvider: FC<AuthProviderProps> = (props) => {
  const { children } = props;
  const [isAuth, setAuth] = useState<boolean>(defaultValues.isAuth);
  const [isLoading, setLoading] = useState<boolean>(defaultValues.isLoading);
  const [userRole, setUserRole] = useState<UserRole>(defaultValues.userRole);
  const [token, setToken] = useState(defaultValues.token);
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
            setToken(token);
            setUserRole(data["role"]);
          } else {
            const isLoggedIn = await refreshTokens();
            if (isLoggedIn) {
              setAuth(true);
              setUserRole(data["role"]);
              setToken(token);
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

  const logout = async () => {
    return fetch(import.meta.env.VITE_API_URL + "/auth/logout", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + token,
      },
      credentials: "include",
    })
      .then(async (res) => {
        if (res.status === 200) {
          setAuth(false);
          localStorage.removeItem("access_token");
          navigate("/auth");
        }
      })
      .catch((err) => console.error(err));
  };

  useEffect(() => {
    const token = localStorage.getItem("access_token");
    if (token)
      return () => {
        getUserInfo(token);
      };

    navigate("/auth");
    return () => {
      setLoading(false);
      setAuth(false);
    };
  }, [getUserInfo, navigate]);

  const exposed = {
    isAuth,
    setAuth,
    isLoading,
    userRole,
    logout,
    token,
  };

  return <Context.Provider value={exposed}>{children}</Context.Provider>;
};

export const useAuth = () => useContext(Context);
