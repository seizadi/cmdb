
CREATE TABLE aws_services (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  region_id int REFERENCES regions(id) ON DELETE CASCADE
);

CREATE TRIGGER aws_services_updated_at
  BEFORE UPDATE OR INSERT ON aws_services
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();
