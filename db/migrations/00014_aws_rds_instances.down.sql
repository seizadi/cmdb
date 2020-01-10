
DROP TRIGGER aws_rds_instances_updated_at on aws_rds_instances;
ALTER TABLE secrets DROP COLUMN aws_rds_instance_id;

DROP TABLE aws_rds_instances;
