CREATE TABLE IF NOT EXISTS cctvs
(
    id         uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    title      VARCHAR(64)     NOT NULL,
    link       VARCHAR(128)     NOT NULL,
    latitude   DECIMAL(10, 8)   NOT NULL,
    longitude  DECIMAL(11, 8)   NOT NULL,
    width      SMALLINT         NOT NULL,
    height     SMALLINT         NOT NULL,
    created_at TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
