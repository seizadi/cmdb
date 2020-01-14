
CREATE TABLE app_versions (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  repo text DEFAULT NULL,
  version text DEFAULT NULL,
  chart_version_id int REFERENCES chart_versions(id) ON DELETE CASCADE,
  application_id int REFERENCES applications(id) ON DELETE CASCADE,
  lifecycle_id int REFERENCES lifecycles(id) ON DELETE CASCADE,
  environment_id int REFERENCES environments(id) ON DELETE CASCADE
);

CREATE TRIGGER app_versions_updated_at
  BEFORE UPDATE OR INSERT ON app_versions
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

