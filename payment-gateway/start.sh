#!/bin/bash

source ./app.env
echo "$DB_SOURCE"
# Check if db-up.sql file exists
if [ ! -f "db/migration/db-up.sql" ]; then
    echo "Error: db-up.sql file not found."
    exit 1
fi

echo "$DB_HOST"
echo "$DB_PORT"
echo "$DB_USER"
echo "$DB_NAME"

# Connect to the PostgreSQL database
psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -w -c "SELECT 1" > /dev/null 2>&1

if [ $? -ne 0 ]; then
    echo "Error: Could not connect to PostgreSQL database."
    exit 1
fi

# Execute the db-up.sql file
psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" < "db/migration/db-up.sql"

if [ $? -ne 0 ]; then
    echo "Error: Could not execute db-up.sql file."
    exit 1
fi

echo "Migrations executed successfully."
