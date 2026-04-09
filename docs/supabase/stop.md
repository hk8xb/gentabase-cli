## supabase-stop

Stops the Gentabase local development stack.

Requires `supabase/config.toml` to be created in your current working directory by running `gentabase init`.

All Docker resources are maintained across restarts.  Use `--no-backup` flag to reset your local development data between restarts.

Use the `--all` flag to stop all local Gentabase projects instances on the machine. Use with caution with `--no-backup` as it will delete all gentabase local projects data.