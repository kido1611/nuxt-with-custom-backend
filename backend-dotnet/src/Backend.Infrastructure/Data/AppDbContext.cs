using Microsoft.EntityFrameworkCore;
using Backend.Domain.Entities;

namespace Backend.Infrastructure.Data;

public class AppDbContext : DbContext
{

    public AppDbContext(DbContextOptions<AppDbContext> options) : base(options)
    {
    }

    public DbSet<User> Users { get; set; }
    public DbSet<Note> Notes { get; set; }
    public DbSet<Session> Sessions { get; set; }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<User>().ToTable("users");
        modelBuilder.Entity<Note>().ToTable("notes");
        modelBuilder.Entity<Session>().ToTable("sessions");

        modelBuilder.Entity<User>()
          .HasIndex(c => c.Email)
          .IsUnique();
        modelBuilder.Entity<User>().Property(b => b.CreatedAt)
          .HasDefaultValueSql("NOW()");
        modelBuilder.Entity<User>().Property(b => b.UpdatedAt)
          .HasDefaultValueSql("NOW()");

        modelBuilder.Entity<Note>().Property(b => b.CreatedAt)
          .HasDefaultValueSql("NOW()");
        modelBuilder.Entity<Note>().Property(b => b.UpdatedAt)
          .HasDefaultValueSql("NOW()");

        modelBuilder.Entity<Session>().Property(b => b.CreatedAt)
          .HasDefaultValueSql("NOW()");
        modelBuilder.Entity<Session>().Property(b => b.UpdatedAt)
          .HasDefaultValueSql("NOW()");

    }
}
