
CREATE TABLE lifecycles (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  config_yaml text DEFAULT NULL,
  lifecycle_id int REFERENCES lifecycles(id) ON DELETE CASCADE
);

CREATE TRIGGER lifecycles_updated_at
  BEFORE UPDATE OR INSERT ON lifecycles
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

