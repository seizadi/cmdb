
CREATE TABLE app_configs (
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  config_yaml text DEFAULT NULL,
  application_id text REFERENCES applications(id) ON DELETE CASCADE,
  lifecycle_id text REFERENCES lifecycles(id) ON DELETE CASCADE,
  environment_id text REFERENCES environments(id) ON DELETE CASCADE
);

CREATE TRIGGER app_configs_updated_at
  BEFORE UPDATE OR INSERT ON app_configs
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

