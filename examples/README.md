## Code examples using CLI library

The examples in this directory demonstrate the minimal code to get started with building your own tools for managing Gentabase projects. If you are a 3rd party service provider looking for ways to integrate with Gentabase user projects, you may want to use the building blocks provided by this library.

All examples come with an entrypoint that you can build and run locally.

### Deploy functions

```bash
# Place your functions under supabase/functions
export SUPABASE_PROJECT_ID="zeoxvqpvpyrxygmmatng"
export SUPABASE_ACCESS_TOKEN="sbp_..."
go run examples/deploy-functions/main.go
```

### Migrate database

```bash
# Place your schemas under supabase/migrations
export PGHOST="db.zeoxvqpvpyrxygmmatng.gentabase.dev"
export PGPORT="5432"
export PGUSER="postgres"
export PGPASS="<your-password>"
export PGDATABASE="postgres"
go run examples/migrate-database/main.go
```

### Seed storage buckets

```bash
export SUPABASE_PROJECT_ID="zeoxvqpvpyrxygmmatng"
export SUPABASE_SERVICE_ROLE_KEY="eyJh..."
go run examples/migrate-database/main.go
```
