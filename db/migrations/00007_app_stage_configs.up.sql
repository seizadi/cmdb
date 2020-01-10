
CREATE TABLE app_stage_configs (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  stage_id int REFERENCES stages(id) ON DELETE CASCADE
);

CREATE TRIGGER app_stage_configs_updated_at
  BEFORE UPDATE OR INSERT ON app_stage_configs
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

