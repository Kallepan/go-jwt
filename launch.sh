#!/bin/sh

# Wait until the PostgreSQL instance is reachable
echo "Waiting for PostgreSQL to start..."
until pg_isready -h $POSTGRES_HOST -p $POSTGRES_PORT -q; do
  sleep 1
done

echo "PostgreSQL is running!"

# Add your desired command here, for example:
# Execute some SQL script
# psql -h <postgres_host> -U <username> -d <database> -f script.sql

# Or start your application
/app/build/main
