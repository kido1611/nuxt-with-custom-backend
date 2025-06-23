CREATE TABLE sessions (
  id VARCHAR(100) NOT NULL PRIMARY KEY,
  user_id VARCHAR,
  csrf_token VARCHAR NOT NULL,
  ip_address VARCHAR(50),
  user_agent VARCHAR(200),
  expired_at DATETIME NOT NULL,
  last_activity_at DATETIME,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_sessions_updated_at
  AFTER UPDATE ON sessions
  FOR EACH ROW
BEGIN
  UPDATE sessions
  SET updated_at = CURRENT_TIMESTAMP
  WHERE id = OLD.id;
END;
