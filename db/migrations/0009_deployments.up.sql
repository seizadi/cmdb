
CREATE TABLE deployments (
  id serial primary key,
  account_id varchar(255),
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  artifact_id int REFERENCES artifacts(id) ON DELETE SET NULL,
  kube_cluster_id int REFERENCES kube_clusters(id) ON DELETE CASCADE
);

CREATE TRIGGER deployments_updated_at
  BEFORE UPDATE OR INSERT ON deployments
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

