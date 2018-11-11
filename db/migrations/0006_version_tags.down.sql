
DROP TRIGGER version_tags_updated_at on version_tags;
ALTER TABLE artifacts DROP COLUMN version_tag_id;

DROP TABLE version_tags;
