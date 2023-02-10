CREATE OR REPLACE FUNCTION TRIGGER_set_updated()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE users (
    id serial,
    username  VARCHAR(80) NOT NULL,
    password VARCHAR(80) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NUll DEFAULT CURRENt_TIMESTAMP
);

CREATE TRIGGER users_set_updated
BEFORE UPDATE ON public.users
FOR EACH ROW
EXECUTE PROCEDURE TRIGGER_set_updated();


insert into users(username, password) values('test', '$2a$04$mWQwJDSYLOEox1QpSU.tpukS68V3M/rXR8RUpyCet4m2xrNuXcyrq');