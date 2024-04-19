#!/bin/bash

source ./app.env
echo "$DB_SOURCE"
# Check if db-up.sql file exists
if [ ! -f "db/migration/db-up.sql" ]; then
    echo "Error: db-up.sql file not found."
    exit 1
fi

echo "Host: $DB_HOST"
echo "port: $DB_PORT"
echo "user: $DB_USER"
echo "db Name: $DB_NAME"
echo "password: $DB_PASSWORD"

# export PWD=$DB_PASSWORD
# psql -U "$DB_USER" -d payment_service -f db/migration/db-up.sql

# Connect to the PostgreSQL database
psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -w -c --password="$DB_PASSWORD" "SELECT 1" > /dev/null 2>&1

if [ $? -ne 0 ]; then
    echo "Error: Could not connect to PostgreSQL database."
    exit 1
fi

# Execute the db-up.sql file
psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -W password="$DB_PASSWORD"< "db/migration/db-up.sql"

if [ $? -ne 0 ]; then
    echo "Error: Could not execute db-up.sql file."
    exit 1
fi

echo "Migrations executed successfully."
