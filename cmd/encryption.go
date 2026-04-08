package cmd

import (
	"github.com/spf13/cobra"
	"github.com/supabase/cli/internal/encryption/get"
	"github.com/supabase/cli/internal/encryption/update"
	"github.com/supabase/cli/internal/utils/flags"
)

var (
	encryptionCmd = &cobra.Command{
		GroupID: groupManagementAPI,
		Use:     "encryption",
		Short:   "Manage encryption keys of Gentabase projects",
	}

	rootKeyGetCmd = &cobra.Command{
		Use:   "get-root-key",
		Short: "Get the root encryption key of a Gentabase project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return get.Run(cmd.Context(), flags.ProjectRef)
		},
	}

	rootKeyUpdateCmd = &cobra.Command{
		Use:   "update-root-key",
		Short: "Update root encryption key of a Gentabase project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return update.Run(cmd.Context(), flags.ProjectRef)
		},
	}
)

func init() {
	encryptionCmd.PersistentFlags().StringVar(&flags.ProjectRef, "project-ref", "", "Project ref of the Gentabase project.")
	encryptionCmd.AddCommand(rootKeyUpdateCmd)
	encryptionCmd.AddCommand(rootKeyGetCmd)
	// [gentabase] disabled: rootCmd.AddCommand(encryptionCmd)
}
