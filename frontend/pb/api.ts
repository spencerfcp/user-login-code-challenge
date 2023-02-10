/* eslint-disable */

export const protobufPackage = "";

export interface User {
  username: string;
}

export interface UserSignupRequest {
  username: string;
  password: string;
}

export interface UserSignupResponse {
  user: User | undefined;
  usernameAlreadyExists: boolean;
  invalidCredentials: boolean;
}

export interface UserLoginRequest {
  username: string;
  password: string;
}

export interface UserLoginResponse {
  user: User | undefined;
  invalidCredentials: boolean;
}

export interface EmptyRequest {
}

export interface EmptyResponse {
}

function createBaseUser(): User {
  return { username: "" };
}

export const User = {
  fromJSON(object: any): User {
    return { username: isSet(object.username) ? String(object.username) : "" };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    return obj;
  },

  create<I extends Exact<DeepPartial<User>, I>>(base?: I): User {
    return User.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.username = object.username ?? "";
    return message;
  },
};

function createBaseUserSignupRequest(): UserSignupRequest {
  return { username: "", password: "" };
}

export const UserSignupRequest = {
  fromJSON(object: any): UserSignupRequest {
    return {
      username: isSet(object.username) ? String(object.username) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: UserSignupRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserSignupRequest>, I>>(base?: I): UserSignupRequest {
    return UserSignupRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserSignupRequest>, I>>(object: I): UserSignupRequest {
    const message = createBaseUserSignupRequest();
    message.username = object.username ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseUserSignupResponse(): UserSignupResponse {
  return { user: undefined, usernameAlreadyExists: false, invalidCredentials: false };
}

export const UserSignupResponse = {
  fromJSON(object: any): UserSignupResponse {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      usernameAlreadyExists: isSet(object.usernameAlreadyExists) ? Boolean(object.usernameAlreadyExists) : false,
      invalidCredentials: isSet(object.invalidCredentials) ? Boolean(object.invalidCredentials) : false,
    };
  },

  toJSON(message: UserSignupResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    message.usernameAlreadyExists !== undefined && (obj.usernameAlreadyExists = message.usernameAlreadyExists);
    message.invalidCredentials !== undefined && (obj.invalidCredentials = message.invalidCredentials);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserSignupResponse>, I>>(base?: I): UserSignupResponse {
    return UserSignupResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserSignupResponse>, I>>(object: I): UserSignupResponse {
    const message = createBaseUserSignupResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.usernameAlreadyExists = object.usernameAlreadyExists ?? false;
    message.invalidCredentials = object.invalidCredentials ?? false;
    return message;
  },
};

function createBaseUserLoginRequest(): UserLoginRequest {
  return { username: "", password: "" };
}

export const UserLoginRequest = {
  fromJSON(object: any): UserLoginRequest {
    return {
      username: isSet(object.username) ? String(object.username) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: UserLoginRequest): unknown {
    const obj: any = {};
    message.username !== undefined && (obj.username = message.username);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserLoginRequest>, I>>(base?: I): UserLoginRequest {
    return UserLoginRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserLoginRequest>, I>>(object: I): UserLoginRequest {
    const message = createBaseUserLoginRequest();
    message.username = object.username ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseUserLoginResponse(): UserLoginResponse {
  return { user: undefined, invalidCredentials: false };
}

export const UserLoginResponse = {
  fromJSON(object: any): UserLoginResponse {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
      invalidCredentials: isSet(object.invalidCredentials) ? Boolean(object.invalidCredentials) : false,
    };
  },

  toJSON(message: UserLoginResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    message.invalidCredentials !== undefined && (obj.invalidCredentials = message.invalidCredentials);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserLoginResponse>, I>>(base?: I): UserLoginResponse {
    return UserLoginResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserLoginResponse>, I>>(object: I): UserLoginResponse {
    const message = createBaseUserLoginResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    message.invalidCredentials = object.invalidCredentials ?? false;
    return message;
  },
};

function createBaseEmptyRequest(): EmptyRequest {
  return {};
}

export const EmptyRequest = {
  fromJSON(_: any): EmptyRequest {
    return {};
  },

  toJSON(_: EmptyRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<EmptyRequest>, I>>(base?: I): EmptyRequest {
    return EmptyRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmptyRequest>, I>>(_: I): EmptyRequest {
    const message = createBaseEmptyRequest();
    return message;
  },
};

function createBaseEmptyResponse(): EmptyResponse {
  return {};
}

export const EmptyResponse = {
  fromJSON(_: any): EmptyResponse {
    return {};
  },

  toJSON(_: EmptyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<EmptyResponse>, I>>(base?: I): EmptyResponse {
    return EmptyResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmptyResponse>, I>>(_: I): EmptyResponse {
    const message = createBaseEmptyResponse();
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
