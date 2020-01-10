
CREATE TABLE aws_rds_instances (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  database_host text DEFAULT NULL,
  database_name text DEFAULT NULL,
  database_user text DEFAULT NULL,
  aws_service_id int REFERENCES aws_services(id) ON DELETE CASCADE
);

CREATE TRIGGER aws_rds_instances_updated_at
  BEFORE UPDATE OR INSERT ON aws_rds_instances
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE secrets ADD COLUMN aws_rds_instance_id int REFERENCES aws_rds_instances(id) ON DELETE CASCADE;
