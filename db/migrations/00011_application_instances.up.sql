
CREATE TABLE application_instances (
   id text primary key,
   account_id text,
   created_at timestamptz DEFAULT current_timestamp,
   updated_at timestamptz DEFAULT NULL,
   name text DEFAULT NULL,
   description text DEFAULT NULL,
   enable bool DEFAULT TRUE,
   config_yaml text DEFAULT NULL,
   chart_version_id text REFERENCES chart_versions(id) ON DELETE CASCADE,
   environment_id text REFERENCES environments(id) ON DELETE CASCADE,
   application_id text REFERENCES applications(id) ON DELETE CASCADE,
   app_version_id text REFERENCES app_versions(id) ON DELETE CASCADE
);

CREATE TRIGGER application_instances_updated_at
    BEFORE UPDATE OR INSERT ON application_instances
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();
