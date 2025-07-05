using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Backend.Infrastructure.Migrations
{
    /// <inheritdoc />
    public partial class AddTriggerUpdate : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.Sql(@"
CREATE TRIGGER update_users_updated_at
  AFTER UPDATE ON users
  FOR EACH ROW
BEGIN
  UPDATE users
  SET updated_at = CURRENT_TIMESTAMP
  WHERE id = NEW.id;
END;
              ");

            migrationBuilder.Sql(@"
CREATE TRIGGER update_notes_updated_at
  AFTER UPDATE ON notes
  FOR EACH ROW
BEGIN
  UPDATE notes
  SET updated_at = CURRENT_TIMESTAMP
  WHERE id = OLD.id;
END;
              ");

            migrationBuilder.Sql(@"
CREATE TRIGGER update_sessions_updated_at
  AFTER UPDATE ON sessions
  FOR EACH ROW
BEGIN
  UPDATE sessions
  SET updated_at = CURRENT_TIMESTAMP
  WHERE id = OLD.id;
END;
                  ");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.Sql(@"
                DROP TRIGGER IF EXISTS update_users_updated_at;
              ");

            migrationBuilder.Sql(@"
                DROP TRIGGER IF EXISTS update_notes_updated_at;
              ");

            migrationBuilder.Sql(@"
                DROP TRIGGER IF EXISTS update_sessions_updated_at;
              ");
        }
    }
}
