
CREATE TABLE chart_versions (
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  repo text DEFAULT NULL,
  version text DEFAULT NULL,
  application_id text REFERENCES applications(id) ON DELETE CASCADE
);

CREATE TRIGGER chart_versions_updated_at
  BEFORE UPDATE OR INSERT ON chart_versions
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

