
DROP TRIGGER aws_services_updated_at on aws_services;
ALTER TABLE aws_rds_instances DROP COLUMN aws_service_id;

DROP TABLE aws_services;
