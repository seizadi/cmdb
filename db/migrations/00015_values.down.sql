
DROP TRIGGER values_updated_at on values;
ALTER TABLE regions DROP COLUMN value_id;
ALTER TABLE stages DROP COLUMN value_id;
ALTER TABLE environments DROP COLUMN value_id;
ALTER TABLE app_region_configs DROP COLUMN value_id;
ALTER TABLE app_stage_configs DROP COLUMN value_id;
ALTER TABLE app_environment_configs DROP COLUMN value_id;
ALTER TABLE application_instances DROP COLUMN value_id;
ALTER TABLE secrets DROP COLUMN value_id;

DROP TABLE values;
