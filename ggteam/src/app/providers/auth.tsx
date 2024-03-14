import {
  Dispatch,
  FC,
  PropsWithChildren,
  SetStateAction,
  createContext,
  useContext,
  useState,
} from "react";

interface AuthContext {
  isAuth: boolean;
  setAuth: Dispatch<SetStateAction<boolean>>;
}

const defaultValues = {
  isAuth: false,
  setAuth: () => null,
};

const Context = createContext(defaultValues as AuthContext);

interface AuthProviderProps extends PropsWithChildren { }

export const AuthProvider: FC<AuthProviderProps> = (props) => {
  const { children } = props;
  const [isAuth, setAuth] = useState<boolean>(defaultValues.isAuth);

  const exposed = {
    isAuth,
    setAuth,
  };

  return <Context.Provider value={exposed}>{children}</Context.Provider>;
};

export const useAuth = () => useContext(Context);
