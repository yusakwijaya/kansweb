#!/bin/bash
set -e

# Debugging: Print the environment variables
echo "POSTGRES_USER: $POSTGRES_USER"
echo "POSTGRES_DB: $POSTGRES_DB"

# Check if the database exists
DB_EXISTS=$(psql -U "$POSTGRES_USER" -d postgres -tAc "SELECT 1 FROM pg_database WHERE datname='$POSTGRES_DB'")
echo "DB_EXISTS: $DB_EXISTS"

if [ "$DB_EXISTS" != "1" ]; then
  echo "Creating database $POSTGRES_DB"
  # Create the database if it does not exist
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres"
      CREATE DATABASE $POSTGRES_DB;
  EOSQL
else
  echo "Database $POSTGRES_DB already exists"
fi

# Drop the public schema if it exists
echo "Dropping existing public schema"
psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "DROP SCHEMA IF EXISTS public CASCADE;"

# Restore the database from the custom format dump
echo "Restoring database $POSTGRES_DB from kansweb.sql"
pg_restore -U "$POSTGRES_USER" -d "$POSTGRES_DB" /docker-entrypoint-initdb.d/kansweb.sql
