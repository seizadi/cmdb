
CREATE TABLE artifacts (
  id serial primary key,
  account_id varchar(255),
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  repo varchar(255) DEFAULT NULL,
  commit varchar(255) DEFAULT NULL
);

CREATE TRIGGER artifacts_updated_at
  BEFORE UPDATE OR INSERT ON artifacts
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

