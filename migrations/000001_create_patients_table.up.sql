CREATE TABLE IF NOT EXISTS patients (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    first_name text NOT NULL,
    last_name text NOT NULL,
    gender CHAR(1) NOT NULL,
    date_of_birth date NOT NULL
);