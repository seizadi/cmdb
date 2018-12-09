
CREATE TABLE manifests (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL,
  repo text DEFAULT NULL,
  commit text DEFAULT NULL,
  values jsonb not null default '{}',
  services jsonb not null default '{}',
  ingress jsonb not null default '{}',
  resources jsonb not null default '{}',
  node_selector jsonb not null default '{}',
  tolerations jsonb not null default '[]',
  affinity jsonb not null default '{}',
  artifact_id int REFERENCES artifacts(id) ON DELETE SET NULL,
  vault_id int REFERENCES vaults(id) ON DELETE SET NULL,
  aws_service_id int REFERENCES aws_services(id) ON DELETE SET NULL
);

CREATE TRIGGER manifests_updated_at
  BEFORE UPDATE OR INSERT ON manifests
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

