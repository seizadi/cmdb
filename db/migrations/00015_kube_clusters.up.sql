
CREATE TABLE kube_clusters (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  network_id int REFERENCES networks(id) ON DELETE CASCADE

);

CREATE TRIGGER kube_clusters_updated_at
  BEFORE UPDATE OR INSERT ON kube_clusters
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

