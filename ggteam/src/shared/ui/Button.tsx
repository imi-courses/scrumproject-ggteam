import { ButtonHTMLAttributes, PropsWithChildren, forwardRef } from "react";
import "styles/Button.css";

interface ButtonProps
  extends ButtonHTMLAttributes<HTMLButtonElement>,
  PropsWithChildren { }

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ children, ...props }, ref) => {
    return (
      <button {...props} ref={ref} className="classButton">
        {children}
      </button>
    );
  },
);
Button.displayName = "Button";

export default Button;
