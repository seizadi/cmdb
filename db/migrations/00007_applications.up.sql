
CREATE TABLE applications (
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  chart text DEFAULT NULL,
  repo text DEFAULT NULL,
  ticket_link text DEFAULT NULL,
  config_yaml text DEFAULT NULL

);

CREATE TRIGGER applications_updated_at
  BEFORE UPDATE OR INSERT ON applications
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

