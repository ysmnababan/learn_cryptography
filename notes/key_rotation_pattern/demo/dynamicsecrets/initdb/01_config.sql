-- ✅ Make sure vault_admin owns the schema
ALTER SCHEMA public OWNER TO vault_admin;

-- ✅ Grant CRUD on existing tables
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO PUBLIC;

-- ✅ Grant sequence usage for existing sequences (needed for SERIAL/BIGSERIAL)
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA public TO PUBLIC;

-- ✅ Default privileges for future tables created by vault_admin
ALTER DEFAULT PRIVILEGES FOR ROLE vault_admin IN SCHEMA public
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO PUBLIC;

-- ✅ Default privileges for future sequences created by vault_admin
ALTER DEFAULT PRIVILEGES FOR ROLE vault_admin IN SCHEMA public
GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO PUBLIC;
