import { postUserSignup, UserSignupError } from "@/api/api";
import Button from "@/components/Button";
import LoginSplash from "@/components/LoginSplash/LoginSplash";
import TextInput from "@/components/TextInput";
import { Form, handleInputChange } from "@/utils/validate/validate";
import Head from "next/head";
import Link from "next/link";
import { UserSignupResponse } from "pb/api";
import { FormEvent, useState } from "react";
import { FaLock, FaUserAlt } from "react-icons/fa";
import { FiAlertCircle } from "react-icons/fi";

export default function SignUp() {
  const initialValue = {
    value: "",
    isValid: null,
  };

  const formValidations: Form = {
    username: initialValue,
    password: initialValue,
  };

  interface FormState {
    error?: UserSignupError;
    isLoading: boolean;
    isFinished?: boolean;
  }

  const [values, setValues] = useState<Form>(formValidations);
  const [formState, setFormState] = useState<FormState>({
    isLoading: false,
    isFinished: false,
  });

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const isValid = values.username.isValid && values.password.isValid;
    if (isValid) {
      setFormState({ ...formState, isLoading: true });
      !formState.isLoading &&
        (await postUserSignup(
          {
            username: values.username.value,
            password: values.password.value,
          },
          (_: UserSignupResponse) => {
            setFormState({ isLoading: false, isFinished: true });
          },
          (err: UserSignupError) => {
            setFormState({ isLoading: false, error: err });

            if (err !== "exists") {
              setValues(formValidations);
            }
          }
        ));
    }
  };

  return (
    <>
      <Head>
        <title>User Login Challenge</title>
        <meta name="description" content="User Signup" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className="flex-grow items-center lg:justify-center flex flex-col">
        <LoginSplash>
          {formState.isFinished ? (
            <span
              data-testid="signup_successful_message"
              className={"text-green-600 font-semibold text-md mt-5"}
            >
              Sign up successful! Click <Link href="/">here </Link> to login.
            </span>
          ) : (
            <>
              <form
                onSubmit={onSubmit}
                className={
                  "flex flex-col gap-2 px-10 mt-5 w-full animate-fadeIn "
                }
                aria-label="Sign Up"
              >
                <TextInput
                  type="text"
                  leftIcon={<FaUserAlt />}
                  name="username"
                  errorText={values.username.error}
                  placeholder="New Username"
                  aria-required="true"
                  data-testid="signup_username"
                  isValid={values.username.isValid && !formState.error}
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
                  errorText={values.password.error}
                  type="password"
                  name="password"
                  placeholder="Password"
                  data-testid="signup_password"
                  aria-required="true"
                  isValid={values.password.isValid && !formState.error}
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
                    {formState.error === "exists" && (
                      <span data-testid="signup_user_exists_error">
                        {"Username already exists"}
                      </span>
                    )}
                    {formState.error === "generic" && (
                      <span data-testid="signup_generic_error">
                        {"There has been an error. Please try again later."}
                      </span>
                    )}
                  </div>
                )}
                <div className={"flex flex-row justify-center mt-6 gap-3"}>
                  <Button
                    aria-label="submit"
                    data-testid="signup_submit_button"
                    loading={formState.isLoading}
                    type="submit"
                    disabled={formState.isLoading}
                  >
                    Sign Up
                  </Button>
                </div>
              </form>
              <Link href="/" className="mt-5 text-xs animate-fadeIn ">
                Already have an account?
              </Link>
            </>
          )}
        </LoginSplash>
      </main>
    </>
  );
}
