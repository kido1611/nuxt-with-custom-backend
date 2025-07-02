-- Add up migration script here
CREATE TABLE notes (
  id TEXT NOT NULL PRIMARY KEY,
  user_id VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  visible_at DATETIME,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME,

  FOREIGN KEY("user_id") REFERENCES users ("id") ON DELETE CASCADE
);

CREATE TRIGGER update_notes_updated_at
  AFTER UPDATE ON notes
  FOR EACH ROW
BEGIN
  UPDATE notes
  SET updated_at = CURRENT_TIMESTAMP
  WHERE id = OLD.id;
END;
