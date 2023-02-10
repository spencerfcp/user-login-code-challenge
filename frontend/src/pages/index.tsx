import { postUserLogin, UserLoginError } from "@/api/api";
import Button from "@/components/Button";
import LoginSplash from "@/components/LoginSplash/LoginSplash";
import TextInput from "@/components/TextInput";
import { Form, handleInputChange } from "@/utils/validate/validate";
import Head from "next/head";
import Link from "next/link";
import { UserLoginResponse } from "pb/api";
import { FormEvent, useState } from "react";
import { FaLock, FaUserAlt } from "react-icons/fa";
import { FiAlertCircle } from "react-icons/fi";

export default function Home() {
  const initialValue = {
    value: "",
    isValid: true,
  };

  const formValidations: Form = {
    username: initialValue,
    password: initialValue,
  };

  interface FormState {
    error?: UserLoginError;
    isLoading: boolean;
    isFinished?: boolean;
  }
  const [values, setValues] = useState<Form>(formValidations);
  const [formState, setFormState] = useState<FormState>({
    isLoading: false,
  });
  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setFormState({ ...formState, isLoading: true });
    !formState.isLoading &&
      (await postUserLogin(
        {
          username: values.username.value,
          password: values.password.value,
        },
        (_: UserLoginResponse) => {
          setFormState({ isFinished: true, isLoading: false });
        },
        (err: UserLoginError) => {
          setFormState({ isLoading: false, error: err });
          setValues({
            username: {
              value: values.username.value,
              isValid: false,
            },
            password: {
              value: "",
              isValid: false,
            },
          });
        }
      ));
  };

  return (
    <>
      <Head>
        <title>User Login Challenge</title>
        <meta name="description" content="User Login Challenge" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className="flex-grow items-center lg:justify-center flex flex-col">
        <LoginSplash>
          {formState.isFinished ? (
            <span
              data-testid="login_successful_message"
              className={"text-green-600 font-semibold text-md mt-5"}
            >
              Login Successful!{" "}
              <a
                onClick={(e) => {
                  e.preventDefault();
                  setValues(formValidations);
                  setFormState({ isLoading: false });
                }}
              >
                Return to home.
              </a>
            </span>
          ) : (
            <>
              <form
                onSubmit={onSubmit}
                className={"flex flex-col w-full gap-2 px-10 mt-5"}
                aria-label="Sign Up"
              >
                <TextInput
                  type="text"
                  leftIcon={<FaUserAlt />}
                  name={"username"}
                  isValid={!formState.error}
                  placeholder="Username"
                  aria-required="true"
                  data-testid="login_username"
                  value={values.username.value}
                  onChange={(e) => {
                    handleInputChange(values, e, "change", setValues);
                  }}
                  onBlur={(e) => {
                    handleInputChange(values, e, "blur", setValues);
                  }}
                />
                <TextInput
                  leftIcon={<FaLock />}
                  isValid={!formState.error}
                  type="password"
                  name={"password"}
                  data-testid="login_password"
                  placeholder="Password"
                  aria-required="true"
                  value={values.password.value}
                  onChange={(e) => {
                    handleInputChange(values, e, "change", setValues);
                  }}
                  onBlur={(e) => {
                    handleInputChange(values, e, "blur", setValues);
                  }}
                />
                {formState.error && (
                  <div
                    className={
                      "flex flex-row justify-end text-sm items-center gap-2 text-red-500"
                    }
                  >
                    <FiAlertCircle />
                    {formState.error === "invalid_creds" && (
                      <span data-testid="login_invalid_creds_error">
                        {"Invalid username or password"}
                      </span>
                    )}
                    {formState.error === "generic" && (
                      <span data-testid="login_generic_error">
                        {"There has been an error. Please try again later."}
                      </span>
                    )}
                  </div>
                )}
                <div className={"flex flex-row justify-center mt-6"}>
                  <Button
                    data-testid="login_submit_button"
                    aria-label="submit"
                    type="submit"
                    disabled={formState.isLoading}
                    loading={formState.isLoading}
                  >
                    {"Sign In"}
                  </Button>
                </div>
              </form>

              <Link href="/signup" className="mt-5 text-xs">
                {"Dont have an account?"}
              </Link>
            </>
          )}
        </LoginSplash>
      </main>
    </>
  );
}
