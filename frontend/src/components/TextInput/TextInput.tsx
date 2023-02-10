import classNames from "classnames";
import { InputHTMLAttributes, useRef } from "react";

interface TextInputProps
  extends Omit<
    InputHTMLAttributes<HTMLInputElement>,
    "type" | "onChange" | "onBlur"
  > {
  label?: string;
  type: "text" | "password";
  errorText?: string;
  isValid?: boolean | null;
  leftIcon?: React.ReactNode;
  onChange?: (e: HTMLInputElement) => void;
  onBlur?: (e: HTMLInputElement) => void;
}
const TextInput = (props: TextInputProps) => {
  const {
    errorText,
    isValid,
    leftIcon,
    type,
    label,
    onChange: onChangeProp,
    onBlur: onBlurProp,
    ...rest
  } = props;

  const inputRef = useRef<HTMLInputElement>(null);
  const onChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    onChangeProp?.(e.currentTarget);

  const onBlur = (e: React.FocusEvent<HTMLInputElement>) => {
    onBlurProp?.(e.currentTarget);
  };

  return (
    <div className="flex flex-col gap-1">
      <span className="flex text-gray-700 text-sm font-bold ">{label}</span>
      <label>
        <div
          className={classNames(
            "flex flex-row border rounded-md  p-2 gap-2 items-center text-gray-500 shadow-inner",
            {
              "border-gray-200": isValid || isValid == null,
              "border-red-200": isValid === false,
            }
          )}
        >
          {leftIcon && leftIcon}
          <input
            ref={inputRef}
            type={type}
            onBlur={onBlur}
            onChange={onChange}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                inputRef.current && onBlurProp?.(e.currentTarget);
              }
            }}
            className="bg-transparent outline-none flex-grow flex"
            {...rest}
          />
        </div>
      </label>
      {errorText && (
        <span className={"flex justify-end text-right text-red-500 text-xs"}>
          {errorText}
        </span>
      )}
    </div>
  );
};

export default TextInput;
