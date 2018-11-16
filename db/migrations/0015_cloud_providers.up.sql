
CREATE TABLE cloud_providers (
  id serial primary key,
  account_id varchar(255),
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  provider int DEFAULT 0,
  account varchar(255) DEFAULT NULL
);

CREATE TRIGGER cloud_providers_updated_at
  BEFORE UPDATE OR INSERT ON cloud_providers
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE regions ADD COLUMN cloud_provider_id int REFERENCES cloud_providers(id) ON DELETE SET NULL;
