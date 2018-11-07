
CREATE TABLE regions (
  id serial primary key,
  account_id varchar(255),
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL
);

CREATE FUNCTION set_updated_at()
  RETURNS trigger as $$
  BEGIN
    NEW.updated_at := current_timestamp;
    RETURN NEW;
  END $$ language plpgsql;

CREATE TRIGGER regions_updated_at
  BEFORE UPDATE OR INSERT ON regions
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();
