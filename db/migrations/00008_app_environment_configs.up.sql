
CREATE TABLE app_environment_configs (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  environment_id int REFERENCES environments(id) ON DELETE CASCADE
);

CREATE TRIGGER app_environment_configs_updated_at
  BEFORE UPDATE OR INSERT ON app_environment_configs
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

