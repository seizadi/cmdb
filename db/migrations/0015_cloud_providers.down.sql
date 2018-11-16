
DROP TRIGGER cloud_providers_updated_at on cloud_providers;
ALTER TABLE regions DROP COLUMN cloud_provider_id;

DROP TABLE cloud_providers;
