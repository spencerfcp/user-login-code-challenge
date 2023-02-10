export interface Form {
  [key: string]: FormValue;
}

export interface FormValue {
  value: string;
  isValid: boolean | null;
  error?: string;
}

type inputChange = "blur" | "change";

const FieldErrors: { [key: string]: string } = {
  username: "Username must be between 2-80 characters",
  password:
    "Password is too weak. Must have a combination of upper and lowercase letters, numbers, and special symbols",
};

export function handleInputChange<T extends Record<string, FormValue>>(
  values: T,
  e: HTMLInputElement,
  type: inputChange,
  setValues: (v: Form) => void
) {
  const { name, value } = e;

  if (type === "blur") {
    const isValid = validateStringField(name, value);
    setValues({
      ...values,
      [name]: {
        value: value,
        isValid: value ? isValid : null,
        error: value && !isValid ? FieldErrors[name] ?? null : null,
      },
    });
  } else {
    setValues({
      ...values,
      [name]: {
        value: value,
        isValid: values[name].isValid != null ? values[name].isValid : null,
      },
    });
  }
}

export const validateStringField = (
  fieldName: string,
  value: string
): boolean => {
  switch (fieldName) {
    case "username":
      return validateUsername(value);
    case "password":
      return validatePassword(value);
    default:
      return false;
  }
};

const validatePassword = (password: string): boolean => {
  return password.match(
    /(?=.{8,})((?=.*\d)(?=.*[a-z])(?=.*[A-Z])|(?=.*\d)(?=.*[a-zA-Z])(?=.*[\W_])|(?=.*[a-z])(?=.*[A-Z])(?=.*[\W_])).*/
  )
    ? true
    : false;
};

export const validateUsername = (name: string): boolean => {
  return name.length > 1 && name.length < 80;
};
