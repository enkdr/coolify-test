# Use the official Postgres image from Docker Hub
FROM postgres:15

# Set environment variables for the database
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydatabase

# Copy SQL scripts to be run at container startup
COPY init-scripts/*.sql /docker-entrypoint-initdb.d/

# Expose the default Postgres port
EXPOSE 5432
