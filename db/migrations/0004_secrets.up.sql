CREATE TABLE secrets
(
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  type text DEFAULT NULL,
  key text DEFAULT NULL
);

CREATE TRIGGER secrets_updated_at
  BEFORE UPDATE OR INSERT ON secrets
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();