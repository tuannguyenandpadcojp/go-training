# Database Schema Creation with Golang-Migrate

This document outlines the steps to create and manage database schemas using golang-migrate.

## Prerequisites
- Install the golang-migrate CLI tool: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## Getting Started
1. Create a new migration folder
```bash
mkdir -p database/migrations
```

2. Create a new migration file to create a new database
```bash
migrate create -ext sql -dir database/migrations -seq create_tm_database
```

3. Edit the newly created migration file

4. Create a new migration file to create a table clients
```bash
migrate create -ext sql -dir database/migrations -seq create_clients_table
```

5. Edit the newly created migration file

6. Create a new migration file to create a table users
```bash
migrate create -ext sql -dir database/migrations -seq create_users_table
```

7. Edit the newly created migration file

8. Create a docker-compose file to run a MySQL database

8. Run docker-compose up to start the MySQL server
```bash
docker-compose up -d
```

8. Run the migration
```bash
make db-migration
```

9. Verify the database and table creation
- Connect to the MySQL server
```bash
make exec-db
```
- Verify the database creation
```sql
SHOW DATABASES;
```
- Verify the tables creation
```sql
SHOW TABLES;
```
- Verify the tables schema
```sql
DESCRIBE clients;
DESCRIBE users;
```
