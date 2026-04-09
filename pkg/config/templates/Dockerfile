# Exposed for updates by .github/dependabot.yml
FROM ghcr.io/hk8xb/postgres:17.6.1.104 AS pg
# Append to ServiceImages when adding new dependencies below
FROM library/kong:2.8.1 AS kong
FROM axllent/mailpit:v1.22.3 AS mailpit
FROM postgrest/postgrest:v14.8 AS postgrest
FROM ghcr.io/hk8xb/postgres-meta:v0.96.2 AS pgmeta
FROM ghcr.io/hk8xb/studio:2026.04.06-sha-b9e83b2 AS studio
FROM darthsim/imgproxy:v3.8.0 AS imgproxy
FROM ghcr.io/hk8xb/edge-runtime:v1.73.3 AS edgeruntime
FROM timberio/vector:0.53.0-alpine AS vector
FROM ghcr.io/hk8xb/supavisor:2.7.4 AS supavisor
FROM ghcr.io/hk8xb/gotrue:v2.188.1 AS gotrue
FROM ghcr.io/hk8xb/realtime:v2.80.12 AS realtime
FROM ghcr.io/hk8xb/storage-api:v1.48.21 AS storage
FROM ghcr.io/hk8xb/logflare:1.36.1 AS logflare
# Append to JobImages when adding new dependencies below
FROM ghcr.io/hk8xb/pgadmin-schema-diff:cli-0.0.5 AS differ
FROM ghcr.io/hk8xb/migra:3.0.1663481299 AS migra
FROM ghcr.io/hk8xb/pg_prove:3.36 AS pgprove
