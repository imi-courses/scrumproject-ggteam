import { Dispatch, FC, SetStateAction } from "react";
import "styles/Input.css"

interface InputProps {
  value: string;
  setValue: Dispatch<SetStateAction<string>>;
}

const Input: FC<InputProps> = (props) => {
  const { value, setValue } = props;
  return <input value={value} onChange={(e) => setValue(e.target.value)} className="classInput"/>;
};

export default Input;
