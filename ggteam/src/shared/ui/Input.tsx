import { Dispatch, FC, SetStateAction } from "react";

interface InputProps {
  value: string;
  setValue: Dispatch<SetStateAction<string>>;
}

const Input: FC<InputProps> = (props) => {
  const { value, setValue } = props;
  return <input value={value} onChange={(e) => setValue(e.target.value)} />;
};

export default Input;
