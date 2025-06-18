ALTER TABLE sessions
ADD COLUMN last_activity_at DATETIME AFTER expired_at;
