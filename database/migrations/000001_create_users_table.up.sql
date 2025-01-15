
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TYPE IF EXISTS user_role;
CREATE TYPE user_role AS ENUM ('admin', 'guest');

CREATE TABLE IF NOT EXISTS users
(
    id             uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email          VARCHAR(100)     NOT NULL UNIQUE,
    name           VARCHAR(100)     NOT NULL,
    role           user_role        NOT NULL DEFAULT 'guest',
    is_active      BOOLEAN          NOT NULL DEFAULT TRUE,
    last_logged_in TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at     TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP
);