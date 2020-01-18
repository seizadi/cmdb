
CREATE TABLE deployments (
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  artifact_id text REFERENCES artifacts(id) ON DELETE SET NULL,
  application_instance_id text REFERENCES application_instances(id) ON DELETE SET NULL,
  kube_cluster_id text REFERENCES kube_clusters(id) ON DELETE CASCADE
);

CREATE TRIGGER deployments_updated_at
  BEFORE UPDATE OR INSERT ON deployments
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

