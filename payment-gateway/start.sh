#!/bin/bash

source ./app.env
# Check if db-up.sql file exists
if [ ! -f "db/migration/db-up.sql" ]; then
    echo "Error: db-up.sql file not found."
    exit 1
fi

echo "starting migration..."

# Connect to the PostgreSQL database
# PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -w -c "SELECT 1" > /dev/null 2>&1
RUN_ON_MYDB="psql -X -U $DB_USER password=$DB_PASSWORD --set ON_ERROR_STOP=on"

if [ $? -ne 0 ]; then
    echo "Error: Could not connect to PostgreSQL database."
    exit 1
fi

# Execute the db-up.sql file
# PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" < "db/migration/db-up.sql"
RUN_ON_MYDB="psql -X -U $DB_USER password=$DB_PASSWORD --set ON_ERROR_STOP=on" < "db/migration/db-up.sql"
if [ $? -ne 0 ]; then
    echo "Error: Could not execute db-up.sql file."
    exit 1
fi

echo "Migrations executed successfully."
