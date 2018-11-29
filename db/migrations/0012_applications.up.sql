
CREATE TABLE applications (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  app_name text DEFAULT NULL,
  repo text DEFAULT NULL,
  version_tag_id int REFERENCES version_tags(id) ON DELETE SET NULL,
  manifest_id int REFERENCES manifests(id) ON DELETE SET NULL
);

CREATE TRIGGER applications_updated_at
  BEFORE UPDATE OR INSERT ON applications
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE containers ADD COLUMN application_id int REFERENCES applications(id) ON DELETE CASCADE;
ALTER TABLE deployments ADD COLUMN application_id int REFERENCES applications(id) ON DELETE CASCADE;
