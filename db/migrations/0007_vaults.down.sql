
DROP TRIGGER vaults_updated_at on vaults;
ALTER TABLE secrets DROP COLUMN vault_id;

DROP TABLE vaults;
