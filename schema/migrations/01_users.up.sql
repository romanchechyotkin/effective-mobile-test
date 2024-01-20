CREATE TYPE gender as ENUM (
    'male',
    'female'
);

CREATE TABLE IF NOT EXISTS users
(
    id bigserial primary key,
    last_name text NOT NULL,
    first_name text NOT NULL,
    second_name text default '',
    age integer NOT NULL,
    gender public.gender NOT NULL,
    nationality text NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);
