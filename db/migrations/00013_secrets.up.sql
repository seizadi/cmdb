CREATE TABLE secrets
(
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  path text DEFAULT NULL,
  type text DEFAULT NULL,
  key text DEFAULT NULL,
  vault_id text REFERENCES vaults(id) ON DELETE CASCADE
);

CREATE TRIGGER secrets_updated_at
  BEFORE UPDATE OR INSERT ON secrets
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();
