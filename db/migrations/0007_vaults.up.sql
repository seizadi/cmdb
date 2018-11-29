
CREATE TABLE vaults
(
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  path text DEFAULT NULL
);

CREATE TRIGGER vaults_updated_at
  BEFORE UPDATE OR INSERT ON vaults
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE secrets ADD COLUMN vault_id int REFERENCES vaults(id) ON DELETE CASCADE;

