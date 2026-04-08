# Gentabase CLI

Manage your Gentabase projects from the command line.

## Installation

```bash
# Recommended — one command, any platform
curl -fsSL https://gentabase.dev/install.sh | sh

# Homebrew (macOS / Linux)
brew install hk8xb/tap/gentabase

# npm
npm install -g @gentabase/cli
```

## Quick Start

```bash
# Authenticate
gentabase login

# Initialize a new project
gentabase init

# Start local development
gentabase start

# Create a migration
gentabase migration new create_todos_table

# Push migrations to your project
gentabase db push

# Generate TypeScript types
gentabase gen types typescript --local > types/gentabase.ts

# Deploy edge functions
gentabase functions deploy my-function

# Stop local containers
gentabase stop
```

## Commands

### Local Development
| Command | Description |
|---|---|
| `gentabase init` | Initialize a local project |
| `gentabase start` | Start containers for local development |
| `gentabase stop` | Stop all local containers |
| `gentabase status` | Show status of local containers |
| `gentabase services` | Show versions of all services |

### Database
| Command | Description |
|---|---|
| `gentabase db push` | Push migrations to remote database |
| `gentabase db pull` | Pull schema from remote database |
| `gentabase db reset` | Reset local database to migrations |
| `gentabase db diff` | Generate a migration by diffing schemas |
| `gentabase db dump` | Dump remote database to SQL |
| `gentabase db lint` | Lint local database |

### Migrations
| Command | Description |
|---|---|
| `gentabase migration new` | Create a new migration |
| `gentabase migration list` | List all migrations |

### Edge Functions
| Command | Description |
|---|---|
| `gentabase functions new` | Create a new edge function |
| `gentabase functions serve` | Serve functions locally |
| `gentabase functions deploy` | Deploy functions to your project |

### Code Generation
| Command | Description |
|---|---|
| `gentabase gen types` | Generate TypeScript types from your schema |

### Authentication
| Command | Description |
|---|---|
| `gentabase login` | Authenticate (opens browser) |
| `gentabase login --token` | Authenticate with a PAT (CI/non-interactive) |
| `gentabase logout` | Log out |
| `gentabase link` | Link to a remote project |

## Authentication

**Interactive (opens browser):**
```bash
gentabase login
```

**Non-interactive (CI/CD):**
```bash
gentabase login --token gbp_<your-personal-access-token>
```

**Environment variable:**
```bash
export GENTABASE_ACCESS_TOKEN=gbp_<your-token>
gentabase projects list
```

Create personal access tokens in the Gentabase dashboard at **Account → Access Tokens**.

## Configuration

Projects are configured via `gentabase/config.toml`. Run `gentabase init` to generate one.

## License

Proprietary — see LICENSE for details.
