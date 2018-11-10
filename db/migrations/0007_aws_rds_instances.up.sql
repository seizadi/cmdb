
CREATE TABLE aws_rds_instances (
  id serial primary key,
  account_id varchar(255),
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL
);

CREATE TRIGGER aws_rds_instances_updated_at
  BEFORE UPDATE OR INSERT ON aws_rds_instances
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

