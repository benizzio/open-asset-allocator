-- clean all schema objects
DO $$
DECLARE
    r RECORD;
BEGIN
    -- Drop all tables
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
            EXECUTE 'DROP TABLE public.' || quote_ident(r.tablename) || ' CASCADE';
    END LOOP;

    -- Drop all sequences
    FOR r IN (SELECT sequencename FROM pg_sequences WHERE schemaname = 'public') LOOP
            EXECUTE 'DROP SEQUENCE public.' || quote_ident(r.sequencename) || ' CASCADE';
    END LOOP;

    -- Drop all views
    FOR r IN (SELECT viewname FROM pg_views WHERE schemaname = 'public') LOOP
            EXECUTE 'DROP VIEW public.' || quote_ident(r.viewname) || ' CASCADE';
    END LOOP;

    -- Drop all functions
    FOR r IN (SELECT routine_name FROM information_schema.routines WHERE specific_schema = 'public') LOOP
            EXECUTE 'DROP FUNCTION public.' || quote_ident(r.routine_name) || ' CASCADE';
    END LOOP;
END $$;

-- grant service user permissions
GRANT USAGE ON SCHEMA 'public' TO open_asset_allocator;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA 'public' TO open_asset_allocator;
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA 'public' TO open_asset_allocator;

ALTER DEFAULT PRIVILEGES IN SCHEMA 'public' GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO open_asset_allocator;
ALTER DEFAULT PRIVILEGES IN SCHEMA 'public' GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO open_asset_allocator;