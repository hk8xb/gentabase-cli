package cmd

import (
	"fmt"
	"os"

	"github.com/go-errors/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/supabase/cli/internal/login"
	"github.com/supabase/cli/internal/utils"
	"golang.org/x/term"
)

var (
	params = login.RunParams{
		OpenBrowser: term.IsTerminal(int(os.Stdin.Fd())),
		Fsys:        afero.NewOsFs(),
	}

	loginCmd = &cobra.Command{
		GroupID: groupLocalDev,
		Use:     "login",
		Short:   "Authenticate using an access token",
		Long: `Authenticate with the Gentabase platform.

Interactive (opens browser):
  gentabase login

Non-interactive (CI / paste token):
  gentabase login --token gbp_<your-token>

Or set the GENTABASE_ACCESS_TOKEN environment variable.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If --token provided, save directly (non-interactive)
			if params.Token != "" {
				if !utils.AccessTokenPattern.MatchString(params.Token) {
					return errors.New(utils.ErrInvalidToken)
				}
				if err := utils.SaveAccessToken(params.Token, params.Fsys); err != nil {
					return err
				}
				fmt.Fprintln(os.Stdout, "Access token saved to profile:", utils.CurrentProfile.Name)
				return nil
			}
			// Try stdin (piped token)
			params.Token = login.ParseAccessToken(os.Stdin)
			if params.Token != "" {
				if !utils.AccessTokenPattern.MatchString(params.Token) {
					return errors.New(utils.ErrInvalidToken)
				}
				if err := utils.SaveAccessToken(params.Token, params.Fsys); err != nil {
					return err
				}
				fmt.Fprintln(os.Stdout, "Access token saved to profile:", utils.CurrentProfile.Name)
				return nil
			}
			// Interactive browser flow
			if !params.OpenBrowser {
				return errors.Errorf(
					"Cannot use automatic login flow inside non-TTY environments. "+
						"Please provide %s flag or set %s.",
					utils.Aqua("--token"), utils.Aqua("GENTABASE_ACCESS_TOKEN"))
			}
			return login.Run(cmd.Context(), os.Stdout, params)
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
	loginFlags.StringVar(&params.Token, "token", "", "Personal access token (PAT) — skip browser flow")
	loginFlags.StringVar(&params.TokenName, "name", "", "Name for the token created via browser flow")
	loginFlags.Lookup("name").DefValue = "auto-generated from hostname"
	loginFlags.Bool("no-browser", false, "Print login URL instead of opening browser")
	rootCmd.AddCommand(loginCmd)
}
