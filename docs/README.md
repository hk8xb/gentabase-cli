## Extended man pages for CLI commands

### Build

Update [version string](https://github.com/hk8xb/gentabase-cli/blob/main/docs/main.go#L33) to match latest release.

```bash
go run docs/main.go > cli_v1_commands.yaml
```

### Release

1. Clone the Gentabase documentation source repository (or your docs site repo)
2. Copy over the CLI reference and reformat using gentabase config

```bash
mv ../cli/cli_v1_commands.yaml specs/
npx prettier -w specs/cli_v1_commands.yaml
```

3. If there are new commands added, update `spec/common-cli-sections.json` in that repository manually
