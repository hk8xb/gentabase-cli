package cmd

import (
	"fmt"
	"os"

	"github.com/go-errors/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/supabase/cli/internal/utils"
)

var (
	loginToken string

	loginCmd = &cobra.Command{
		GroupID: groupLocalDev,
		Use:     "login",
		Short:   "Authenticate using an access token",
		Long: `Authenticate with a personal access token (PAT).

Create a PAT in the Gentabase dashboard at Account → Access Tokens,
then run:

  gentabase login --token gbp_<your-token>

Alternatively, set the GENTABASE_ACCESS_TOKEN environment variable.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if loginToken == "" {
				return errors.New("Please provide a token with --token flag or set GENTABASE_ACCESS_TOKEN")
			}
			// Validate format
			if !utils.AccessTokenPattern.MatchString(loginToken) {
				return errors.New(utils.ErrInvalidToken)
			}
			// Save token
			if err := utils.SaveAccessToken(loginToken, afero.NewOsFs()); err != nil {
				return err
			}
			fmt.Fprintln(os.Stdout, "Access token saved to profile:", utils.CurrentProfile.Name)
			return nil
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			if prof := viper.GetString("PROFILE"); viper.IsSet("PROFILE") {
				return utils.SaveProfileName(prof, afero.NewOsFs())
			}
			return nil
		},
	}
)

func init() {
	loginFlags := loginCmd.Flags()
	loginFlags.StringVar(&loginToken, "token", "", "Personal access token (PAT) from Gentabase dashboard")
	loginFlags.StringP("name", "n", "", "Name for this token in local settings")
	rootCmd.AddCommand(loginCmd)
}
