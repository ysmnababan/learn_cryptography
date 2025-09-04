CREATE ROLE "{{name}}" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}';

-- allow login
GRANT CONNECT ON DATABASE mydb TO "{{name}}";

-- allow schema access
GRANT USAGE ON SCHEMA public TO "{{name}}";

-- allow CRUD on existing tables (created before this user existed)
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO "{{name}}";
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA public TO "{{name}}";