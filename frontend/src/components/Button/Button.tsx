import { ButtonHTMLAttributes } from "react";
import classNames from "classnames";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  loading?: boolean;
}
const Button = (props: ButtonProps) => {
  const { className, children, loading, ...rest } = props;
  return (
    <button
      disabled={rest.disabled ?? false}
      className={classNames("py-1.5 px-5 rounded-md flex flex-row", className, {
        "bg-gray-500": rest.disabled,
        "bg-blue-500": !rest.disabled,
      })}
      {...rest}
    >
      {loading ? (
        <div className="animate-spin inline-block w-6 h-6 border-[3px] border-current border-t-transparent text-white rounded-full" />
      ) : (
        children
      )}
    </button>
  );
};

export default Button;
