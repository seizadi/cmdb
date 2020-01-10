CREATE TABLE app_region_configs (
    id serial primary key,
    account_id text,
    created_at timestamptz DEFAULT current_timestamp,
    updated_at timestamptz DEFAULT NULL,
    name text DEFAULT NULL,
    description text DEFAULT NULL,
    region_id int REFERENCES regions(id) ON DELETE CASCADE
);

CREATE TRIGGER app_region_configs_updated_at
    BEFORE UPDATE OR INSERT ON app_region_configs
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();
