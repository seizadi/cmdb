CREATE TABLE version_tags
(
  id serial primary key,
  account_id varchar(255),
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  version varchar(255) DEFAULT NULL,
  repo varchar(255) DEFAULT NULL,
  commit varchar(255) DEFAULT NULL
);

CREATE TRIGGER version_tags_updated_at
  BEFORE UPDATE OR INSERT ON version_tags
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE artifacts ADD COLUMN version_tag_id int REFERENCES version_tags(id) ON DELETE SET NULL;
