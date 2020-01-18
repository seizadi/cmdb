
CREATE TABLE environments (
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  config_yaml text DEFAULT NULL,
  lifecycle_id text REFERENCES lifecycles(id) ON DELETE CASCADE
);

CREATE TRIGGER environments_updated_at
  BEFORE UPDATE OR INSERT ON environments
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();
