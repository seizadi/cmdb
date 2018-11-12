
DROP TRIGGER regions_updated_at on regions;
ALTER TABLE environments DROP COLUMN region_id;

DROP TABLE regions;
