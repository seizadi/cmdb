
DROP TRIGGER applications_updated_at on applications;
ALTER TABLE containers DROP COLUMN application_id;
ALTER TABLE deployments DROP COLUMN application_id;
DROP TABLE applications;
