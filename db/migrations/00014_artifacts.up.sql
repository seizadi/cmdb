
CREATE TABLE artifacts (
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  repo text DEFAULT NULL,
  commit text DEFAULT NULL,
  chart_version_id text REFERENCES chart_versions(id) ON DELETE CASCADE
);

CREATE TRIGGER artifacts_updated_at
  BEFORE UPDATE OR INSERT ON artifacts
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

