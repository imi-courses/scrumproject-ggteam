import { ButtonHTMLAttributes, PropsWithChildren, forwardRef } from "react";

interface ButtonProps
  extends ButtonHTMLAttributes<HTMLButtonElement>,
  PropsWithChildren { }

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ children, ...props }, ref) => {
    return (
      <button {...props} ref={ref}>
        {children}
      </button>
    );
  },
);
Button.displayName = "Button";

export default Button;
