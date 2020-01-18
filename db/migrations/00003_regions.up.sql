
CREATE TABLE regions (
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  cloud_provider_id text REFERENCES cloud_providers(id) ON DELETE CASCADE

);

CREATE TRIGGER regions_updated_at
  BEFORE UPDATE OR INSERT ON regions
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();
