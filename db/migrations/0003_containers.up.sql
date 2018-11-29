
CREATE TABLE containers (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  container_name text DEFAULT NULL,
  image_repo text DEFAULT NULL,
  image_tag text DEFAULT NULL,
  image_pull_policy text DEFAULT NULL,
  digest text DEFAULT NULL
);

CREATE TRIGGER containers_updated_at
  BEFORE UPDATE OR INSERT ON containers
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

