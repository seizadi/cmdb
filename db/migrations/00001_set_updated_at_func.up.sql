
CREATE FUNCTION set_updated_at()
  RETURNS trigger as $$
  BEGIN
    NEW.updated_at := current_timestamp;
    RETURN NEW;
  END $$ language plpgsql;

