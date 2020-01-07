
CREATE TABLE values (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  keys jsonb not null default '{}',
  aws_rds_instance_id int REFERENCES aws_rds_instances(id) ON DELETE CASCADE
);

CREATE TRIGGER values_updated_at
  BEFORE UPDATE OR INSERT ON values
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE regions ADD COLUMN value_id int REFERENCES values(id) ON DELETE SET NULL;
ALTER TABLE stages ADD COLUMN value_id int REFERENCES values(id) ON DELETE SET NULL;
ALTER TABLE environments ADD COLUMN value_id int REFERENCES values(id) ON DELETE SET NULL;
ALTER TABLE applications ADD COLUMN value_id int REFERENCES values(id) ON DELETE SET NULL;
ALTER TABLE application_instances ADD COLUMN value_id int REFERENCES values(id) ON DELETE SET NULL;
ALTER TABLE secrets ADD COLUMN value_id int REFERENCES values(id) ON DELETE SET NULL;
