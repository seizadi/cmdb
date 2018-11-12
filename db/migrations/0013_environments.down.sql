
DROP TRIGGER environments_updated_at on environments;
ALTER TABLE applications DROP COLUMN environment_id;

DROP TABLE environments;
