-- +migrate Up
CREATE TYPE company_type AS ENUM ('corporations', 'non_profit', 'cooperative', 'sole_proprietorship');

CREATE TABLE IF NOT EXISTS company
(
    id SERIAL PRIMARY KEY NOT NULL
    ,name VARCHAR(15) NOT NULL
    ,description VARCHAR(3000) NOT NULL
    ,employees_amount INT NOT NULL
    ,registered BOOLEAN NOT NULL
    ,type company_type 
);

CREATE UNIQUE INDEX IF NOT EXISTS company_name ON company (name);
