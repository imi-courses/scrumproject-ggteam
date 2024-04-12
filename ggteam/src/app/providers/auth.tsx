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

type UserRole = "admin" | "employee" | undefined;

interface AuthContext {
  isAuth: boolean;
  setAuth: Dispatch<SetStateAction<boolean>>;
  isLoading: boolean;
  userRole: UserRole;
  logout: () => void;
  token: string;
  setToken: Dispatch<SetStateAction<string>>;
  setUserRole: Dispatch<SetStateAction<UserRole>>;
  getUserInfo: () => void;
}

const defaultValues: AuthContext = {
  isAuth: false,
  setAuth: () => null,
  isLoading: true,
  userRole: undefined,
  logout: () => null,
  token: "",
  setToken: () => null,
  setUserRole: () => null,
  getUserInfo: () => null,
};

const Context = createContext(defaultValues);

interface AuthProviderProps extends PropsWithChildren { }

export const AuthProvider: FC<AuthProviderProps> = (props) => {
  const { children } = props;
  const [isAuth, setAuth] = useState<boolean>(defaultValues.isAuth);
  const [isLoading, setLoading] = useState<boolean>(defaultValues.isLoading);
  const [userRole, setUserRole] = useState<UserRole>(defaultValues.userRole);
  const [token, setToken] = useState(defaultValues.token);

  const refreshTokens = useCallback(async () => {
    return fetch(import.meta.env.VITE_API_URL + "/auth/refresh", {
      method: "GET",
      credentials: "include",
    })
      .then(async (res) => {
        const data = await res.json();
        if (res.status === 200) {
          setToken(data["access_token"]);
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

  const getUserInfo = async (t?: string) => {
    let tt = "";
    if (t) tt = t;
    else tt = token;
    const response = await fetch(import.meta.env.VITE_API_URL + "/auth/me", {
      method: "GET",
      headers: {
        Authorization: "Bearer " + tt,
      },
      credentials: "include",
    });

    const data = await response.json();
    if (response.status === 200) {
      setAuth(true);
      setToken(tt);
      setUserRole(data["role"]);
    } else {
      const isLoggedIn = await refreshTokens();
      setAuth(isLoggedIn);
      if (isLoggedIn) setUserRole(data["role"]);
    }
    setLoading(false);
  };

  const getUserInfoCallback = useCallback(getUserInfo, [refreshTokens, token]);

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
        } else if (await refreshTokens()) {
          setAuth(false);
          localStorage.removeItem("access_token");
        }
      })
      .catch((err) => console.error(err));
  };

  useEffect(() => {
    const token = localStorage.getItem("access_token");
    if (token) {
      return () => {
        getUserInfoCallback(token);
      };
    }

    return () => {
      setLoading(false);
      setAuth(false);
    };
  }, [getUserInfoCallback]);

  const exposed = {
    isAuth,
    setAuth,
    isLoading,
    userRole,
    logout,
    token,
    setToken,
    setUserRole,
    getUserInfo,
  };

  return <Context.Provider value={exposed}>{children}</Context.Provider>;
};

export const useAuth = () => useContext(Context);
