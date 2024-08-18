CREATE DATABASE mydatabase;

\c mydatabase  -- Connect to the newly created database

-- Your custom SQL commands
ALTER USER authenticator WITH PASSWORD :'pgpass';
ALTER USER pgbouncer WITH PASSWORD :'pgpass';
ALTER USER myuser WITH PASSWORD :'pgpass';


CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT
);

INSERT INTO items (name, description) VALUES 
('Item 1', 'Description of Item 1'),
('Item 2', 'Description of Item 2');
