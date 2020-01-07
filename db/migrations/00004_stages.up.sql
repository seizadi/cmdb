
CREATE TABLE stages (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  type int DEFAULT 0,
  region_id int REFERENCES regions(id) ON DELETE CASCADE

);

CREATE TRIGGER stages_updated_at
  BEFORE UPDATE OR INSERT ON stages
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

