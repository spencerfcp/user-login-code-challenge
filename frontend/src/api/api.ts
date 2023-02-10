import axios from "axios";
import {
  UserLoginRequest,
  UserLoginResponse,
  UserSignupRequest,
  UserSignupResponse,
} from "pb/api";

const goApi = process.env.NEXT_PUBLIC_API_HOST;

export type UserExistsError = "exists";
export type GenericError = "generic";
export type InvalidCredentials = "invalid_creds";

export type StandardApiError = GenericError | undefined;
export type UserLoginError = InvalidCredentials | StandardApiError;
export type UserSignupError = UserExistsError | StandardApiError;

export const postUserLogin = async (
  data: UserLoginRequest,
  onSuccess?: (resp: UserLoginResponse) => void,
  onError?: (error?: UserLoginError) => void
) => {
  try {
    const resp = await axios.post<UserLoginResponse>(`${goApi}/login`, data, {
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json;charset=UTF-8",
      },
    });

    if (!resp.data) {
      throw new Error("Empty response data");
    }
    if (resp.data.invalidCredentials) {
      onError && onError("invalid_creds");
      return;
    }

    onSuccess && onSuccess(resp.data);
  } catch (_) {
    onError && onError("generic");
  }
};

export const postUserSignup = async (
  data: UserSignupRequest,
  onSuccess?: (resp: UserSignupResponse) => void,
  onError?: (error?: UserSignupError) => void
) => {
  try {
    const resp = await axios.post<UserSignupResponse>(`${goApi}/user`, data, {
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json;charset=UTF-8",
      },
    });
    if (!resp.data) {
      throw new Error("Empty response data");
    }

    if (resp.data?.usernameAlreadyExists) {
      onError && onError("exists");
      return;
    }
    onSuccess && onSuccess(resp.data);
  } catch (_) {
    onError && onError("generic");
  }
};
