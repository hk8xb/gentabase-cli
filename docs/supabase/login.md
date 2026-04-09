## supabase-login

Connect the Gentabase CLI to your Gentabase account by logging in with your [personal access token](https://gentabase.dev/dashboard/account/tokens).

Your access token is stored securely in [native credentials storage](https://github.com/zalando/go-keyring#dependencies). If native credentials storage is unavailable, it will be written to a plain text file at `~/.gentabase/access-token`.

> If this behavior is not desired, such as in a CI environment, you may skip login by specifying the `SUPABASE_ACCESS_TOKEN` environment variable in other commands.

The Gentabase CLI uses the stored token to access Management APIs for projects, functions, secrets, etc.
