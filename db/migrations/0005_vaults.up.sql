CREATE TABLE vaults
(
  id serial primary key,
  account_id varchar(255),
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  path varchar(255) DEFAULT NULL
);

CREATE TRIGGER vaults_updated_at
  BEFORE UPDATE OR INSERT ON vaults
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();