
CREATE TABLE environments (
  id serial primary key,
  account_id varchar(255),
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  code int DEFAULT 0
);

CREATE TRIGGER environments_updated_at
  BEFORE UPDATE OR INSERT ON environments
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE applications ADD COLUMN environment_id int REFERENCES environments(id) ON DELETE SET NULL;
