
CREATE TABLE application_instances (
   id serial primary key,
   account_id text,
   created_at timestamptz DEFAULT current_timestamp,
   updated_at timestamptz DEFAULT NULL,
   name text DEFAULT NULL,
   description text DEFAULT NULL,
   app_version_id int REFERENCES app_versions(id) ON DELETE CASCADE,
   environment_id int REFERENCES environments(id) ON DELETE CASCADE,
   application_id int REFERENCES applications(id) ON DELETE CASCADE
);

CREATE TRIGGER application_instances_updated_at
    BEFORE UPDATE OR INSERT ON application_instances
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();
