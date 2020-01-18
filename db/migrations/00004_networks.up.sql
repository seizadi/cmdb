
CREATE TABLE networks (
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  region_id text REFERENCES regions(id) ON DELETE CASCADE

);

CREATE TRIGGER networks_updated_at
  BEFORE UPDATE OR INSERT ON networks
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

