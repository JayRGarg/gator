# gator

A powerful CLI tool for managing and automating your development workflow.

## Prerequisites

Before you begin, ensure you have the following installed:
- [Go](https://golang.org/doc/install) (version 1.21 or higher)
- [PostgreSQL](https://www.postgresql.org/download/) (version 15 or higher)

## Installation

Install the gator CLI tool using Go:

```bash
go install github.com/jayrgarg/gator@latest
```

## Configuration

1. Create a configuration file at `~/.gator/config.yaml` with the following structure:

```yaml
database:
  host: localhost
  port: 5432
  user: your_username
  password: your_password
  dbname: your_database
```

2. Ensure your PostgreSQL server is running and the database is created.

## Usage

The gator CLI provides several commands to help manage your development workflow:

```bash
# Initialize a new project
gator init

# Create a new migration
gator migration create

# Apply pending migrations
gator migration up

# Rollback the last migration
gator migration down

# Show migration status
gator migration status
```

For more information about available commands, run:

```bash
gator --help
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
