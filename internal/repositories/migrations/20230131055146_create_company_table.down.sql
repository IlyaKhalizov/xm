-- +migrate Down
DROP INDEX IF EXISTS company_name;
DROP TABLE IF EXISTS company;
DROP TYPE company_type;