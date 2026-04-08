package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/supabase/cli/internal/secrets/list"
	"github.com/supabase/cli/internal/secrets/set"
	"github.com/supabase/cli/internal/secrets/unset"
	"github.com/supabase/cli/internal/utils/flags"
)

var (
	secretsCmd = &cobra.Command{
		GroupID: groupManagementAPI,
		Use:     "secrets",
		Short:   "Manage Gentabase secrets",
	}

	secretsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all secrets on Gentabase",
		Long:  "List all secrets in the linked project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return list.Run(cmd.Context(), flags.ProjectRef, afero.NewOsFs())
		},
	}

	secretsSetCmd = &cobra.Command{
		Use:   "set <NAME=VALUE> ...",
		Short: "Set a secret(s) on Gentabase",
		Long:  "Set a secret(s) to the linked Gentabase project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return set.Run(cmd.Context(), flags.ProjectRef, envFilePath, args, afero.NewOsFs())
		},
	}

	secretsUnsetCmd = &cobra.Command{
		Use:   "unset [NAME] ...",
		Short: "Unset a secret(s) on Gentabase",
		Long:  "Unset a secret(s) from the linked Gentabase project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return unset.Run(cmd.Context(), flags.ProjectRef, args, afero.NewOsFs())
		},
	}
)

func init() {
	secretsCmd.PersistentFlags().StringVar(&flags.ProjectRef, "project-ref", "", "Project ref of the Gentabase project.")
	secretsSetCmd.Flags().StringVar(&envFilePath, "env-file", "", "Read secrets from a .env file.")
	secretsCmd.AddCommand(secretsListCmd)
	secretsCmd.AddCommand(secretsSetCmd)
	secretsCmd.AddCommand(secretsUnsetCmd)
	rootCmd.AddCommand(secretsCmd)
}
